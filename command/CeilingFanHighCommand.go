package main

type CeilingFanHighCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int
}

func NewCeilingFanHighCommand(ceilingFan *CeilingFan) Command {
	return &CeilingFanHighCommand{
		ceilingFan: ceilingFan,
	}
}

func (cc *CeilingFanHighCommand) execute() {
	cc.prevSpeed = cc.ceilingFan.getSpeed()
	cc.ceilingFan.high()
}

func (cc *CeilingFanHighCommand) undo() {
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

func (cc *CeilingFanHighCommand) getName() string {
	return "CeilingFanHighCommand"
}
