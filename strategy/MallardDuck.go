package main

import "fmt"

type MallardDuck struct {
	fly   FlyBehavior
	quack QuackBehavior

	NormalDuck
}

func NewMallardDuck() Duck {
	Mallard := MallardDuck{
		fly:   NewFlyWithWings(),
		quack: NewQuack(),
	}
	return &Mallard
}

func (m *MallardDuck) Display() {
	fmt.Println("I'm a real Mallard duck")
}

func (m *MallardDuck) Fly() {
	m.fly.fly()
}

func (m *MallardDuck) Quack() {
	m.quack.quack()
}
