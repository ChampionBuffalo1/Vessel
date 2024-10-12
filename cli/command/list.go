package command

import (
	"fmt"

	"github.com/ChampionBuffalo1/vessel/internal"
	"github.com/ChampionBuffalo1/vessel/internal/command/image"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all images in the local store",
	Long:  `usage: vessel list`,
	Run: func(cmd *cobra.Command, args []string) {
		listImages()
	},
}

func listImages() {

	client, ctx, err := internal.NewContainerdClient()
	if err != nil {
		fmt.Println("Error creating containerd client:", err)
		return
	}

	err = image.List(client, ctx)
	if err != nil {
		fmt.Printf("Error listing images: %v\n", err)
	}
}
