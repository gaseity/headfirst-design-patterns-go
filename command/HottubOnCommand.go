package main

type HottubOnCommand struct {
	hottub *Hottub
}

func NewHottubOnCommand(hottub *Hottub) Command {
	return &HottubOnCommand{
		hottub: hottub,
	}
}

func (hc *HottubOnCommand) execute() {
	hc.hottub.on()
	hc.hottub.setTemperature(104)
	hc.hottub.circulate()
}

func (hc *HottubOnCommand) undo() {
	hc.hottub.off()
}

func (cc *HottubOnCommand) getName() string {
	return "HottubOnCommand"
}
