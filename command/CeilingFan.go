package main

import "fmt"

type CeilingFan struct {
	HIGH   int
	MEDIUM int
	LOW    int
	OFF    int

	location string
	speed    int
}

func NewCeilingFan(location string) *CeilingFan {
	return &CeilingFan{
		HIGH:   3,
		MEDIUM: 2,
		LOW:    1,
		OFF:    0,

		location: location,
	}
}

func (cf *CeilingFan) high() {
	// turns the ceiling fan on to high
	cf.speed = cf.HIGH
	fmt.Println(cf.location + " ceiling fan is on high")
}

func (cf *CeilingFan) medium() {
	// turns the ceiling fan on to medium
	cf.speed = cf.MEDIUM
	fmt.Println(cf.location + " ceiling fan is on medium")
}

func (cf *CeilingFan) low() {
	// turns the ceiling fan on to low
	cf.speed = cf.LOW
	fmt.Println(cf.location + " ceiling fan is on low")
}

func (cf *CeilingFan) off() {
	// turns the ceiling fan off
	cf.speed = cf.OFF
	fmt.Println(cf.location + " ceiling fan is off")
}

func (cf *CeilingFan) getSpeed() int {
	return cf.speed
}
