package main

import "fmt"

type TV struct {
	location string
	channel  int
}

func NewTV(location string) *TV {
	return &TV{
		location: location,
	}
}

func (t *TV) on() {
	fmt.Println(t.location + " TV is on")
}

func (t *TV) off() {
	fmt.Println(t.location + " TV is off")
}

func (t *TV) setInputChannel() {
	t.channel = 3
	fmt.Println(t.location + " TV channel is set for DVD")
}
