package main

import (
	"fmt"
	"strconv"
)

type Hottub struct {
	level       bool
	temperature int
}

func NewHottub() *Hottub {
	return &Hottub{}
}

func (ht *Hottub) on() {
	ht.level = true
}

func (ht *Hottub) off() {
	ht.level = false
}

func (ht *Hottub) circulate() {
	if ht.level {
		fmt.Println("Hottub is bubbling!")
	}
}

func (ht *Hottub) jetsOn() {
	if ht.level {
		fmt.Println("Hottub jets are on")
	}
}

func (ht *Hottub) jetsOff() {
	if ht.level {
		fmt.Println("Hottub jets are off")
	}
}

func (ht *Hottub) setTemperature(temperature int) {
	if temperature > ht.temperature {
		fmt.Println("Hottub is heating to a steaming " + strconv.Itoa(ht.temperature) + " degrees")
	} else {
		fmt.Println("Hottub is cooling to " + strconv.Itoa(ht.temperature) + " degrees")
	}
	ht.temperature = temperature
}
