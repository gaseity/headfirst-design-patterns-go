package main

type CeilingFanOffCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int
}

func NewCeilingFanOffCommand(ceilingFan *CeilingFan) Command {
	return &CeilingFanOffCommand{ceilingFan: ceilingFan}
}

func (cc *CeilingFanOffCommand) execute() {
	cc.prevSpeed = cc.ceilingFan.getSpeed()
	cc.ceilingFan.off()
}

func (cc *CeilingFanOffCommand) undo() {
	switch cc.prevSpeed {
	case cc.ceilingFan.HIGH:
		cc.ceilingFan.high()

	case cc.ceilingFan.MEDIUM:
		cc.ceilingFan.medium()

	case cc.ceilingFan.LOW:
		cc.ceilingFan.low()

	default:
		cc.ceilingFan.off()
	}
}

func (cc *CeilingFanOffCommand) getName() string {
	return "CeilingFanOffCommand"
}
