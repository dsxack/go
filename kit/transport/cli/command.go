package cli

import "context"

type kitEndpoint = func(ctx context.Context, request interface{}) (response interface{}, err error)

type Command struct{}

type Commands []Command

func NewCommandGroup(name string, commands ...Command) Command {
	return Command{}
}

func NewCommand(name string, endpoint kitEndpoint) Command {
	return Command{}
}
