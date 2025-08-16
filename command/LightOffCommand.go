package main

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) Command {
	return &LightOffCommand{
		light: light,
	}
}

func (lc *LightOffCommand) execute() {
	lc.light.off()
}

func (lc *LightOffCommand) undo() {
	lc.light.on()
}

func (cc *LightOffCommand) getName() string {
	return "LightOffCommand"
}
