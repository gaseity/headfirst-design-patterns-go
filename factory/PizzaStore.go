package main

import "fmt"

type PizzaStore interface {
	orderPizza(string) Pizza
}

type PizzaStoreSoul struct {
}

func (pss *PizzaStoreSoul) cookPizza(pizza Pizza) {
	fmt.Println("--- Making a " + pizza.getName() + " ---")
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
}
