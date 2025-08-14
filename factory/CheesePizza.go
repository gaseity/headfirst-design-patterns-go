package main

import "fmt"

type CheesePizza struct {
	ingredientFactory PizzaIngredientFactory

	PizzaSoul
}

func NewCheesePizza(ingredientFactory PizzaIngredientFactory) Pizza {
	p := CheesePizza{
		ingredientFactory: ingredientFactory,
	}
	return &p
}

func (cp *CheesePizza) prepare() {
	fmt.Println("Preparing " + cp.name)
	cp.dough = cp.ingredientFactory.createDough()
	cp.sauce = cp.ingredientFactory.createSauce()
	cp.cheese = cp.ingredientFactory.createCheese()
}
