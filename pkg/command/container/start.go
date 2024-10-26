package container

import (
	"context"
	"errors"
	"log/slog"
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
		slog.Error("Error loading container", "error", err)
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
				slog.Error("Error creating task", "error", err)
				return err
			}
		} else {
			slog.Error("Error getting task", "error", err)
			return err
		}
	}
	defer func() {
		_, err := task.Delete(ctx)
		if err != nil {
			slog.Error("Error deleting task", "error", err)
		}
	}() // Always delete the task as to not leave container in stopped state

	// Set up a channel for waiting on task to exit
	exitChannel, err := task.Wait(ctx)
	if err != nil {
		slog.Error("Error waiting for task", "error", err)
		return err
	}

	status, err := task.Status(ctx)
	if err != nil {
		slog.Error("Error getting task status", "error", err)
		return err
	}

	if status.Status == containerd.Stopped {
		return errors.New(`stopped container cannot be re-started. please make sure the container task was properly removed`)
	} else if status.Status == containerd.Pausing {
		return errors.New("please wait for the container to finish pause before resuming it")
	} else if status.Status == containerd.Paused {
		// Resume the container task
		if err := task.Resume(ctx); err != nil {
			slog.Error("Error resuming task", "error", err)
			return err
		}
	} else if status.Status == containerd.Created {
		// Run the container task
		if err := task.Start(ctx); err != nil {
			slog.Error("Error starting task", "error", err)
			return err
		}
	}

	// Set up a handler to wait for the Ctrl + C input
	// once we get the signal we terminate the task with SIGTERM (man 7 signal)
	interruptC := make(chan os.Signal, 1)
	signal.Notify(interruptC, syscall.SIGINT)
	<-interruptC

	// Get the task if it is still running
	runningTask, err := container.Task(ctx, nil)
	if err != nil && errdefs.IsNotFound(err) {
		// No task found for the container
		return nil
	}

	if err := runningTask.Kill(ctx, syscall.SIGTERM); err != nil {
		slog.Error("Failed in sending sigterm", "error", err)
		return err
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, pkg.ContainerStopTimeout*time.Second)
	defer cancel()
	select {
	case <-timeoutCtx.Done():
		if err := runningTask.Kill(ctx, syscall.SIGKILL); err != nil {
			slog.Error("Failure in sending sigkill")
			return err
		}
		status, err := runningTask.Delete(ctx)
		if err != nil {
			slog.Error("Failure in deleting task", "error", err)
			return err
		}
		slog.Error("Task deleted", "status", status)
	case exitCode := <-exitChannel:
		code, _, err := exitCode.Result()
		if err != nil {
			slog.Error("Failure in getting exit code")
			return err
		}
		slog.Error("Task exited", "code", code)
	}
	return nil
}
