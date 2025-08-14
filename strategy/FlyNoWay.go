package main

import "fmt"

type flyNoWay struct {
}

func NewFlyNoWay() FlyBehavior {
	return &flyNoWay{}
}

func (f *flyNoWay) fly() error {
	fmt.Println("I can't fly")

	return nil
}
