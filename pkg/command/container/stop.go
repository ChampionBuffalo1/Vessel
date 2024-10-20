package container

import (
	"context"
	"errors"
	"fmt"
	"syscall"
	"time"

	"github.com/ChampionBuffalo1/vessel/pkg"
	containerd "github.com/containerd/containerd/v2/client"
)

func Stop(client *containerd.Client, ctx context.Context, img string) error {
	container, err := client.LoadContainer(ctx, ContainerID)
	if err != nil {
		fmt.Println("Error loading container")
		return err
	}
	task, err := container.Task(ctx, nil)
	if err != nil {
		fmt.Println("Error getting task:")
		return err
	}

	status, err := task.Status(ctx)
	if err != nil {
		fmt.Println("Error getting task status")
		return err
	}

	if status.Status != containerd.Running && status.Status != containerd.Paused {
		return errors.New("container is not running")
	}

	exitC, err := task.Wait(ctx)
	if err != nil {
		fmt.Println("Error on waiting for task")
		return err
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(pkg.ContainerStopTimeout*time.Second))
	defer cancel()

	if err = task.Kill(ctx, syscall.SIGTERM); err != nil {
		fmt.Println("Error stopping container:", err)
		return err
	}

	select {
	case <-timeoutCtx.Done():
		if err := task.Kill(ctx, syscall.SIGKILL); err != nil {
			fmt.Println("Failure in sending sigkill")
			return err
		}
		status, err := task.Delete(ctx)
		if err != nil {
			fmt.Println("Failure in deleting task")
			return err
		}
		fmt.Println("Task deleted", status)
	case exitCode := <-exitC:
		code, _, err := exitCode.Result()
		if err != nil {
			fmt.Println("Failure in getting exit code")
			return err
		}
		fmt.Println("Task exit with status code: ", code)
	}

	fmt.Println("Container stopped")
	return nil
}
