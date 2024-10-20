package command

import (
	"github.com/ChampionBuffalo1/vessel/cli/commands/container"
	"github.com/ChampionBuffalo1/vessel/cli/commands/image"

	"github.com/spf13/cobra"
)

func AttachCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(image.NewImageCommand())
	rootCmd.AddCommand(container.NewContainerCommand())
}
