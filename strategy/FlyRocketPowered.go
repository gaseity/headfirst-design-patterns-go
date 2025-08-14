package main

import "fmt"

type flyRocketPowered struct {
}

func NewFlyRocketPowered() FlyBehavior {
	return &flyRocketPowered{}
}

func (f *flyRocketPowered) fly() error {
	fmt.Println("I'm flying with a rocket")

	return nil
}
