package container

import (
	"context"
	"errors"
	"fmt"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/errdefs"
)

func Remove(client *containerd.Client, ctx context.Context, containerID string) error {
	container, err := client.LoadContainer(ctx, containerID)
	if err != nil {
		return err
	}
	task, err := container.Task(ctx, nil)
	if err != nil && !errdefs.IsNotFound(err) {
		return err
	}
	status, err := task.Status(ctx)
	if err != nil {
		if status.Status == containerd.Running {
			return errors.New("cannot remove a running container")
		}
		return err
	}

	err = container.Delete(ctx, containerd.WithSnapshotCleanup)

	if err != nil {
		fmt.Println("Error deleting container", err)
		return errors.New("cannot remove container")
	}

	return nil
}
