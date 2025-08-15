package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type ChocolateBoiler struct {
	empty  bool
	boiled bool
}

var uniqueInstance *ChocolateBoiler

func newChocolateBoiler() {
	fmt.Println("Creating unique instance of Chocolate Boiler")

	uniqueInstance = &ChocolateBoiler{
		empty:  true,
		boiled: false,
	}

}

func GetInstance() *ChocolateBoiler {
	if uniqueInstance == nil {
		once.Do(newChocolateBoiler)
	}
	fmt.Printf("Returning instance of Chocolate Boiler %p\n", uniqueInstance)
	return uniqueInstance
}

func (cb *ChocolateBoiler) fill() {
	if cb.isEmpty() {
		cb.empty = false
		cb.boiled = false
		// fill the boiler with a milk/chocolate mixture
	}
	fmt.Println("fill:", cb)
}

func (cb *ChocolateBoiler) drain() {
	if !cb.isEmpty() && cb.isBoiled() {
		// drain the boiled milk and chocolate
		cb.empty = true
	}
	fmt.Println("ddrain:", cb)
}

func (cb *ChocolateBoiler) boil() {
	if !cb.isEmpty() && !cb.isBoiled() {
		// bring the contents to a boil
		cb.boiled = true
	}
	fmt.Println("boil:", cb)
}

func (cb *ChocolateBoiler) isEmpty() bool {
	return cb.empty
}

func (cb *ChocolateBoiler) isBoiled() bool {
	return cb.boiled
}
