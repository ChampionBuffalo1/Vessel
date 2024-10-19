package command

import "github.com/spf13/cobra"

func AttachCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(listCmd, pullCmd)
}
