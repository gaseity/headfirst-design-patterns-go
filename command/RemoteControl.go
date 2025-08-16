package main

import (
	"strconv"
	"strings"
)

// This is the invoker
type RemoteControl struct {
	onCommands  [7]Command
	offCommands [7]Command
	undoCommand Command
}

func NewRemoteControl() *RemoteControl {
	rc := RemoteControl{}

	noCommand := NewNoCommand()
	for i := 0; i < 7; i++ {
		rc.onCommands[i] = noCommand
		rc.offCommands[i] = noCommand
	}
	rc.undoCommand = noCommand

	return &rc
}

func (rc *RemoteControl) setCommand(slot int, onCommand Command, offCommand Command) {
	rc.onCommands[slot] = onCommand
	rc.offCommands[slot] = offCommand
}

func (rc *RemoteControl) onButtonWasPushed(slot int) {
	rc.onCommands[slot].execute()
	rc.undoCommand = rc.onCommands[slot]
}

func (rc *RemoteControl) offButtonWasPushed(slot int) {
	rc.offCommands[slot].execute()
	rc.undoCommand = rc.offCommands[slot]
}

func (rc *RemoteControl) undoButtonWasPushed() {
	rc.undoCommand.undo()
}

func (rc *RemoteControl) toString() string {
	var result strings.Builder

	result.WriteString("\n------ Remote Control -------\n")
	for i, _ := range rc.onCommands {
		result.WriteString("[slot " + strconv.Itoa(i) + "] " + rc.onCommands[i].getName() + "    " + rc.offCommands[i].getName() + "\n")
	}
	result.WriteString("[undo] " + rc.undoCommand.getName() + "\n")
	return result.String()
}
