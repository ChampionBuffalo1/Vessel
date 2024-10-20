package command

import (
	"github.com/ChampionBuffalo1/vessel/cli/commands/container"
	"github.com/spf13/cobra"
)

func AttachCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(listCmd, pullCmd)
	rootCmd.AddCommand(container.NewContainerCommand())
}
