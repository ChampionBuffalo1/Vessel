package image

import (
	"fmt"

	"github.com/ChampionBuffalo1/vessel/pkg"
	"github.com/ChampionBuffalo1/vessel/pkg/command/image"

	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull an image from a registry",
	Long:  `usage: vessel pull [OPTIONS] NAME[:TAG|@DIGEST]`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// TODO: add charm cli UI here
		} else {
			pullImage(args[0])
		}
	},
}

func pullImage(imageName string) {
	client, ctx, err := pkg.NewContainerdClient()
	if err != nil {
		fmt.Println("Error creating containerd client:", err)
		return
	}
	err = image.Pull(client, ctx, imageName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Image pulled successfully")
}
