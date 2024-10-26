package container

import (
	"fmt"

	"github.com/ChampionBuffalo1/vessel/pkg"
	"github.com/ChampionBuffalo1/vessel/pkg/command/container"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [container ID]",
	Short: "Remove a container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		removeContainer(args[0])
	},
}

func removeContainer(containerID string) {
	client, ctx, err := pkg.NewContainerdClient()
	if err != nil {
		fmt.Println("Error creating containerd client:", err)
		return
	}
	err = container.Remove(client, ctx, containerID)
	if err != nil {
		fmt.Printf("Error removing container: %v\n", err)
	}
}
