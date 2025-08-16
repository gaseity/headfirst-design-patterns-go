package main

type TVOffCommand struct {
	tv *TV
}

func NewTVOffCommand(tv *TV) Command {
	return &TVOffCommand{
		tv: tv,
	}
}

func (t *TVOffCommand) execute() {
	t.tv.off()
}

func (t *TVOffCommand) undo() {
	t.tv.on()
}

func (cc *TVOffCommand) getName() string {
	return "TVOffCommand"
}
