package main

type TVOnCommand struct {
	tv *TV
}

func NewTVOnCommand(tv *TV) Command {
	return &TVOnCommand{
		tv: tv,
	}
}

func (t *TVOnCommand) execute() {
	t.tv.on()
	t.tv.setInputChannel()
}

func (t *TVOnCommand) undo() {
	t.tv.off()
}

func (cc *TVOnCommand) getName() string {
	return "TVOnCommand"
}
