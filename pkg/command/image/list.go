package image

import (
	"context"
	"fmt"

	containerd "github.com/containerd/containerd/v2/client"
)

func List(client *containerd.Client, ctx context.Context) error {
	images, err := client.ListImages(ctx)
	if err != nil {
		return err
	}
	// TODO: use charm CLI to render a table
	fmt.Println("Images:")
	for _, image := range images {
		fmt.Println(image.Name())
	}
	return nil
}
