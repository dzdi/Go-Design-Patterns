package command21

import "fmt"

type ICommand interface {
	Execute() error
}

type StartCommand struct{}

func (c *StartCommand) Execute() error {
	fmt.Println("game start")
	return nil
}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

type ArchiveCommand struct{}

func (a *ArchiveCommand) Execute() error {
	fmt.Println("game archive")
	return nil
}

func NewArchiveCommand() *ArchiveCommand {
	return &ArchiveCommand{}
}
