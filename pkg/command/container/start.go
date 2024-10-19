package container

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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
	defer task.Delete(ctx)
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
		fmt.Println("Failure in sending sigkill", err)
		return err
	}

	exitCodeStatus := <-exitChannel
	code, _, err := exitCodeStatus.Result()
	if err != nil {
		return err
	}
	fmt.Printf("Container Task exit with code: %d", code)
	return nil
}
