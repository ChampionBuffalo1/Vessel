package container

import (
	"fmt"

	"github.com/ChampionBuffalo1/vessel/pkg"
	"github.com/ChampionBuffalo1/vessel/pkg/command/container"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [CONTAINER_ID]",
	Short: "Starts a container with the matching ID",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		containerID := ""
		if len(args) > 0 {
			containerID = args[0]
		} else {
			containerID = "my-container-random-id"
		}
		startContainer(containerID)
	},
}

func startContainer(containerID string) {
	client, ctx, err := pkg.NewContainerdClient()
	if err != nil {
		fmt.Println("Error creating containerd client:", err)
		return
	}
	err = container.Start(client, ctx, containerID)
	if err != nil {
		fmt.Println("Error starting the container: ", err)
	}
}
