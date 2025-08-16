package main

import "fmt"

type NoCommand struct {
}

func NewNoCommand() Command {
	return &NoCommand{}
}
func (nc *NoCommand) execute() {
	fmt.Println("NoCommand" + " execute")
}
func (nc *NoCommand) undo() {
	fmt.Println("NoCommand" + " undo")
}

func (cc *NoCommand) getName() string {
	return "NoCommand"
}
