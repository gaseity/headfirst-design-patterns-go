package main

type StereoOffCommand struct {
	stereo *Stereo
}

func NewStereoOffCommand(stereo *Stereo) Command {
	return &StereoOffCommand{
		stereo: stereo,
	}
}

func (sc *StereoOffCommand) execute() {
	sc.stereo.off()
}

func (sc *StereoOffCommand) undo() {
	sc.stereo.on()
}

func (cc *StereoOffCommand) getName() string {
	return "StereoOffCommand"
}
