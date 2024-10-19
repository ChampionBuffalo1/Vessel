package image

import (
	"context"
	"errors"
	"fmt"

	containerd "github.com/containerd/containerd/v2/client"
)

func Pull(client *containerd.Client, ctx context.Context, img string) error {
	img = fmt.Sprintf("docker.io/library/%s", img)
	// Check for the image in the local store
	_, err := client.GetImage(ctx, img)
	if err != nil {
		// Pull the image from the registry into containerd store
		_, err = client.Pull(ctx, img, containerd.WithPullUnpack)
		if err != nil {
			return errors.New("pull image failed: " + err.Error())
		}
	}
	return nil
}
