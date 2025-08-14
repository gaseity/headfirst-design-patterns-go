package main

import "fmt"

type flyWithWings struct {
}

func NewFlyWithWings() FlyBehavior {
	return &flyWithWings{}
}

func (f *flyWithWings) fly() error {
	fmt.Println("I'm flying!!")

	return nil
}
