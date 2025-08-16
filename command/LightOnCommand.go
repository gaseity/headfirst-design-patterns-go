package main

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) Command {
	return &LightOnCommand{
		light: light,
	}
}

func (lc *LightOnCommand) execute() {
	lc.light.on()
}

func (lc *LightOnCommand) undo() {
	lc.light.off()
}

func (cc *LightOnCommand) getName() string {
	return "LightOnCommand"
}
