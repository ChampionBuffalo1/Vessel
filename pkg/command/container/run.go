package container

import (
	"context"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/pkg/oci"
)

const (
	ContainerID = "my-container-random-id"
	snapshotID  = "my-snapshot-random-id"
)

func Run(client *containerd.Client, ctx context.Context, imgName string) error {
	img, err := client.GetImage(ctx, imgName)
	if err != nil {
		return err
	}
	container, err := client.LoadContainer(ctx, ContainerID)
	// Create container if not found
	if err != nil {
		container, err = client.NewContainer(
			ctx,
			ContainerID,
			containerd.WithImage(img),
			containerd.WithNewSnapshot(snapshotID, img),
			containerd.WithNewSpec(oci.WithImageConfig(img)),
		)
		if err != nil {
			return err
		}
	}
	return Start(client, ctx, container.ID())
}
