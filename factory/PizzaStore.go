package main

import "fmt"

type PizzaStore interface {
	orderPizza(string) Pizza
	bake()
	cut()
	box()
}

type PizzaStoreSoul struct {
}

func (pss *PizzaStoreSoul) bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (pss *PizzaStoreSoul) cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (pss *PizzaStoreSoul) box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (pss *PizzaStoreSoul) makingPizza(pizza Pizza) {
	fmt.Println("---Soul Making a " + pizza.getName() + " ---")
	pizza.prepare()
	pss.bake()
	pss.cut()
	pss.box()
}
