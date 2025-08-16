package main

import (
	"fmt"
)

type Light struct {
	location string
	level    int
}

func NewLight(location string) *Light {
	return &Light{
		location: location,
	}
}

func (l *Light) on() {
	l.level = 100
	fmt.Println("Light is on")
}

func (l *Light) off() {
	l.level = 0
	fmt.Println("Light is off")
}
