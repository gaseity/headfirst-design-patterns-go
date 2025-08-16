package main

import "fmt"

func main() {

	remoteControl := NewRemoteControl()

	light := NewLight("Living Room")
	tv := NewTV("Living Room")
	stereo := NewStereo("Living Room")
	hottub := NewHottub()

	lightOn := NewLightOnCommand(light)
	stereoOn := NewStereoOnCommand(stereo)
	tvOn := NewTVOnCommand(tv)
	hottubOn := NewHottubOnCommand(hottub)
	lightOff := NewLightOffCommand(light)
	stereoOff := NewStereoOffCommand(stereo)

	tvOff := NewTVOffCommand(tv)
	hottubOff := NewHottubOffCommand(hottub)

	partyOn := []Command{lightOn, stereoOn, tvOn, hottubOn}
	partyOff := []Command{lightOff, stereoOff, tvOff, hottubOff}

	partyOnMacro := NewMacroCommand(partyOn)
	partyOffMacro := NewMacroCommand(partyOff)

	remoteControl.setCommand(0, partyOnMacro, partyOffMacro)
	remoteControl.setCommand(2, lightOn, lightOff)
	remoteControl.setCommand(3, stereoOn, stereoOff)

	fmt.Println(remoteControl.toString())
	fmt.Println("--- Pushing Macro On---")
	remoteControl.onButtonWasPushed(0)
	fmt.Println("--- Pushing Macro Off---")
	remoteControl.offButtonWasPushed(0)

	fmt.Println("--- Pushing Macro 1 2 3---")
	remoteControl.onButtonWasPushed(1)
	remoteControl.onButtonWasPushed(2)
	remoteControl.onButtonWasPushed(3)

	fmt.Println("--- Pushing Macro undo---")
	remoteControl.undoButtonWasPushed()
}
