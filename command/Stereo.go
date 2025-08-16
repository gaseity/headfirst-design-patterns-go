package main

import (
	"fmt"
	"strconv"
)

type Stereo struct {
	location string
}

func NewStereo(location string) *Stereo {
	return &Stereo{
		location: location,
	}
}

func (s *Stereo) on() {
	fmt.Println(s.location + " stereo is on")
}

func (s *Stereo) off() {
	fmt.Println(s.location + " stereo is off")
}

func (s *Stereo) setCD() {
	fmt.Println(s.location + " stereo is set for CD input")
}

func (s *Stereo) setDVD() {
	fmt.Println(s.location + " stereo is set for DVD input")
}

func (s *Stereo) setRadio() {
	fmt.Println(s.location + " stereo is set for Radio")
}

func (s *Stereo) setVolume(volume int) {
	// code to set the volume
	// valid range: 1-11 (after all 11 is better than 10, right?)
	fmt.Println(s.location + " Stereo volume set to " + strconv.Itoa(volume))
}
