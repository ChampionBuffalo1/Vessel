package command

import (
	"github.com/ChampionBuffalo1/vessel/cli/commands/container"
	"github.com/ChampionBuffalo1/vessel/cli/commands/image"

	"github.com/spf13/cobra"
)

func GetCommands() []*cobra.Command {
	commands := []*cobra.Command{
		pullCmd,
		image.NewImageCommand(),
		container.NewContainerCommand(),
	}
	return commands
}
