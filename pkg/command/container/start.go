package container

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ChampionBuffalo1/vessel/pkg"
	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/pkg/cio"
)

func Start(client *containerd.Client, ctx context.Context, containerID string) error {
	container, err := client.LoadContainer(ctx, containerID)
	if err != nil {
		return err
	}
	// Create a runtime task
	// cio.WithStdio will attach the task's stdio to the current process's stdio
	// The task has only been created within the container and not started
	task, err := container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
	if err != nil {
		return err
	}
	// Set up a channel for waiting on task to exit
	exitChannel, err := task.Wait(ctx)
	if err != nil {
		return err
	}
	// Run the container task
	if err := task.Start(ctx); err != nil {
		return err
	}

	// Setup  a handler to wait for the Ctrl + C input
	// once we get the signal we terminal the task with sigterm
	interruptC := make(chan os.Signal, 1)
	signal.Notify(interruptC, syscall.SIGINT)
	<-interruptC

	if err := task.Kill(ctx, syscall.SIGTERM); err != nil {
		fmt.Println("Failed in sending sigterm")
		return err
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(pkg.ContainerStopTimeout*time.Second))
	defer cancel()
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
	case exitCode := <-exitChannel:
		code, _, err := exitCode.Result()
		if err != nil {
			fmt.Println("Failure in getting exit code")
			return err
		}
		fmt.Println("Task exit with status code: ", code)
	}
	return nil
}
