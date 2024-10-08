package cli

import (
	"os"

	"github.com/ChampionBuffalo1/vessel/cli/command"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vessel",
	Short: "Vessel is a simple wrapper around containerd",
	Long:  `usage: vessel [OPTIONS] COMMAND`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(command.PullCmd)
	rootCmd.AddCommand(command.ListCmd)
}
