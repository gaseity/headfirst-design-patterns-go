package main

import "fmt"

type Duck interface {
	Display()
	Fly()
	Quack()
	Swim()
}

func NewDuck(name string) Duck {

	if name == "Mallard" {
		return NewMallardDuck()
	} else if name == "Model" {
		return NewModelDuck()
	}

	return nil
}

type NormalDuck struct {
}

func (m *NormalDuck) Fly() {
	fmt.Println("Fly?")
}

func (m *NormalDuck) Quack() {
	fmt.Println("Quack?")
}
func (m *NormalDuck) Swim() {
	fmt.Println("All ducks float, even decoys!")
}
