package main

type HottubOffCommand struct {
	hottub *Hottub
}

func NewHottubOffCommand(hottub *Hottub) Command {
	return &HottubOffCommand{
		hottub: hottub,
	}
}

func (hc *HottubOffCommand) execute() {
	hc.hottub.setTemperature(98)
	hc.hottub.off()
}

func (hc *HottubOffCommand) undo() {
	hc.hottub.on()
}

func (cc *HottubOffCommand) getName() string {
	return "HottubOffCommand"
}
