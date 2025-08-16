package main

type MacroCommand struct {
	commands []Command
}

func NewMacroCommand(commands []Command) Command {
	return &MacroCommand{
		commands: commands,
	}
}

func (mc *MacroCommand) execute() {
	for _, cmd := range mc.commands {
		cmd.execute()
	}
}

func (mc *MacroCommand) undo() {
	for _, cmd := range mc.commands {
		cmd.undo()
	}
}

func (mc *MacroCommand) getName() string {
	return "MacroCommand"
}
