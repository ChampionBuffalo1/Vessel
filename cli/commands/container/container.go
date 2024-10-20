package container

import "github.com/spf13/cobra"

func NewContainerCommand() *cobra.Command {
	containerCmd := &cobra.Command{
		Use:   "container",
		Short: "Container commands",
	}
	containerCmd.AddCommand(runCmd)
	return containerCmd
}
