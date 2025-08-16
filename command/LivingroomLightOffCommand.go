package main

type LivingroomLightOffCommand struct {
	light *Light
}

func NewLivingroomLightOffCommand(light *Light) Command {
	return &LivingroomLightOffCommand{
		light: light,
	}
}

func (lc *LivingroomLightOffCommand) execute() {
	lc.light.off()
}

func (lc *LivingroomLightOffCommand) undo() {
	lc.light.on()
}

func (cc *LivingroomLightOffCommand) getName() string {
	return "LivingroomLightOffCommand"
}
