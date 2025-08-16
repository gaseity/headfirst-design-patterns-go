package main

type StereoOnCommand struct {
	stereo *Stereo
}

func NewStereoOnCommand(stereo *Stereo) Command {
	return &StereoOnCommand{
		stereo: stereo,
	}
}

func (sc *StereoOnCommand) execute() {
	sc.stereo.on()
}

func (sc *StereoOnCommand) undo() {
	sc.stereo.off()
}

func (cc *StereoOnCommand) getName() string {
	return "StereoOnCommand"
}
