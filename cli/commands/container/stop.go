package container

import (
	"fmt"

	"github.com/ChampionBuffalo1/vessel/pkg"
	"github.com/ChampionBuffalo1/vessel/pkg/command/container"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop CONTAINER_ID",
	Short: "Stop a running container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stopContainer(args[0])
	},
}

func stopContainer(containerID string) {
	client, ctx, err := pkg.NewContainerdClient()
	if err != nil {
		fmt.Println("Error creating containerd client:", err)
		return
	}
	err = container.Stop(client, ctx, containerID)
	if err != nil {
		fmt.Printf("Error stopping the container: %v\n", err)
	}
}
