package main

import "fmt"

type ModelDuck struct {
	fly   FlyBehavior
	quack QuackBehavior

	NormalDuck
}

func NewModelDuck() *ModelDuck {
	return &ModelDuck{

		fly:   NewFlyNoWay(),
		quack: NewQuack(),
	}
}

func (m *ModelDuck) Display() {
	fmt.Println("I'm a model duck")
}

func (m *ModelDuck) Fly() {
	m.fly.fly()
}

func (m *ModelDuck) Quack() {
	m.quack.quack()
}
