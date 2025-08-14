package main

import "fmt"

type FakeQuack struct {
}

func NewFakeQuack() *FakeQuack {
	return &FakeQuack{}
}

func (q *FakeQuack) quack() {
	fmt.Println("Qwak")
}
