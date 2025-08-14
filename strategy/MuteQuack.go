package main

import "fmt"

type MuteQuack struct {
}

func NewMuteQuack() *MuteQuack {
	return &MuteQuack{}
}

func (q *MuteQuack) quack() {
	fmt.Println("<< Silence >>")
}
