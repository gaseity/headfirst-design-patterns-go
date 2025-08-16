package main

type LivingroomLightOnCommand struct {
	light *Light
}

func NewLivingroomLightOnCommand(light *Light) Command {
	return &LivingroomLightOnCommand{
		light: light,
	}
}

func (lc *LivingroomLightOnCommand) execute() {
	lc.light.on()
}

func (lc *LivingroomLightOnCommand) undo() {
	lc.light.off()
}

func (cc *LivingroomLightOnCommand) getName() string {
	return "LivingroomLightOnCommand"
}
