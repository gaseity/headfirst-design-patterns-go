package main

type CeilingFanMediumCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int
}

func NewCeilingFanMediumCommand(ceilingFan *CeilingFan) Command {
	return &CeilingFanMediumCommand{
		ceilingFan: ceilingFan,
	}
}

func (cc *CeilingFanMediumCommand) execute() {
	cc.prevSpeed = cc.ceilingFan.getSpeed()
	cc.ceilingFan.medium()
}

func (cc *CeilingFanMediumCommand) undo() {
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

func (cc *CeilingFanMediumCommand) getName() string {
	return "CeilingFanMediumCommand"
}
