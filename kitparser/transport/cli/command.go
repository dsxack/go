package cli

import (
	kitcli "github.com/dsxack/go/v2/kit/transport/cli"
)

type Commands kitcli.Commands

func NewCommands(commands Commands) kitcli.Commands {
	return kitcli.Commands(commands)
}
