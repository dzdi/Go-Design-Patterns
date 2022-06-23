package command21

import "fmt"

type Command func() error

func StartCommandFunc() Command {
	return func() error {
		fmt.Println("game start")
		return nil
	}
}

func ArchiveCommandFunc() Command {
	return func() error {
		fmt.Println("game archive")
		return nil
	}
}
