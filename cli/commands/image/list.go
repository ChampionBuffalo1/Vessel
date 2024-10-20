package image

import (
	"fmt"

	"github.com/ChampionBuffalo1/vessel/pkg"
	"github.com/ChampionBuffalo1/vessel/pkg/command/image"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all images in the local store",
	Long:  `usage: vessel list`,
	Run: func(cmd *cobra.Command, args []string) {
		listImages()
	},
}

func listImages() {
	client, ctx, err := pkg.NewContainerdClient()
	if err != nil {
		fmt.Println("Error creating containerd client:", err)
		return
	}
	err = image.List(client, ctx)
	if err != nil {
		fmt.Printf("Error listing images: %v\n", err)
	}
}
