package cli

import (
	"os"

	command "github.com/ChampionBuffalo1/vessel/cli/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vessel",
	Short: "Vessel is a simple wrapper around containerd",
	Long:  `usage: vessel [OPTIONS] COMMAND`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	command.AttachCommands(rootCmd)
}
