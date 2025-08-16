package main

type StereoOnWithCDCommand struct {
	stereo *Stereo
}

func NewStereoOnWithCDCommand(stereo *Stereo) Command {
	return &StereoOnWithCDCommand{
		stereo: stereo,
	}
}

func (sc *StereoOnWithCDCommand) execute() {
	sc.stereo.on()
	sc.stereo.setCD()
	sc.stereo.setVolume(11)
}

func (sc *StereoOnWithCDCommand) undo() {
	sc.stereo.off()
}

func (cc *StereoOnWithCDCommand) getName() string {
	return "StereoOnWithCDCommand"
}
