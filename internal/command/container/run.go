package container

import (
	"context"
	"fmt"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/pkg/oci"
)

const (
	containerID = "my-container-random-id"
	snapshotID  = "my-snapshot-random-id"
)

func Run(client *containerd.Client, ctx context.Context, imgName string) error {
	img, err := client.GetImage(ctx, imgName)
	if err != nil {
		return err
	}
	containers, err := client.Containers(ctx, "id=="+containerID)
	// Create container if not found
	if err != nil || len(containers) == 0 {
		_, err = client.NewContainer(
			ctx,
			containerID,
			containerd.WithNewSnapshot(snapshotID, img),
			containerd.WithNewSpec(oci.WithImageConfig(img)),
		)
		if err != nil {
			return err
		}
	}
	// Start the container
	err = Start(client, ctx, containerID)
	fmt.Println("Error from starting in client: ", err)

	return nil
}
