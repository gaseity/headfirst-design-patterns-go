package main

import "fmt"

type Quack struct {
}

func NewQuack() *Quack {
	return &Quack{}
}

func (q *Quack) quack() {
	fmt.Println("Quack")
}
