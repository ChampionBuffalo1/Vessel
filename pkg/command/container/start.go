package container

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ChampionBuffalo1/vessel/pkg"
	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/pkg/cio"
	"github.com/containerd/errdefs"
)

func Start(client *containerd.Client, ctx context.Context, containerID string) error {
	container, err := client.LoadContainer(ctx, containerID)
	if err != nil {
		fmt.Println("Error loading container")
		return err
	}

	task, err := container.Task(ctx, nil)
	if err != nil {
		if errdefs.IsNotFound(err) {
			// Create a runtime task
			// cio.WithStdio will attach the task's stdio to the current process's stdio
			// The task has only been created within the container and not started
			task, err = container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
			if err != nil {
				fmt.Println("Error creating task")
				return err
			}
		} else {
			fmt.Println("Error getting task")
			return err
		}
	}
	defer task.Delete(ctx) // Always delete the task as to not leave container in stopped state

	// Set up a channel for waiting on task to exit
	exitChannel, err := task.Wait(ctx)
	if err != nil {
		fmt.Println("Error waiting for task")
		return err
	}

	status, err := task.Status(ctx)
	if err != nil {
		fmt.Println("Error getting task status")
		return err
	}

	if status.Status == containerd.Stopped {
		return fmt.Errorf(`stopped container cannot be re-started. please make sure the container task is deleted.
		You can use the command "ctr -n vessel task rm %s" to delete the task`, containerID)
	} else if status.Status == containerd.Pausing {
		return errors.New("please wait for the container to finish pause before resuming it")
	} else if status.Status == containerd.Paused {
		// Resume the container task
		if err := task.Resume(ctx); err != nil {
			fmt.Println("Error resuming task")
			return err
		}
	} else if status.Status == containerd.Created {
		// Run the container task
		if err := task.Start(ctx); err != nil {
			fmt.Println("Error starting task")
			return err
		}
	}

	// Setup  a handler to wait for the Ctrl + C input
	// once we get the signal we terminal the task with sigterm
	interruptC := make(chan os.Signal, 1)
	signal.Notify(interruptC, syscall.SIGINT)
	<-interruptC

	// TODO: Handle the case where a user can run `stop` command before SIGINT is received:w
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
