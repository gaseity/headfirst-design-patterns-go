package main

import "fmt"

type ClamPizza struct {
	ingredientFactory PizzaIngredientFactory

	PizzaSoul
}

func NewClamPizza(ingredientFactory PizzaIngredientFactory) Pizza {
	cp := ClamPizza{
		ingredientFactory: ingredientFactory,
	}
	return &cp
}

func (cp *ClamPizza) prepare() {
	fmt.Println("Preparing " + cp.name)
	cp.dough = cp.ingredientFactory.createDough()
	cp.sauce = cp.ingredientFactory.createSauce()
	cp.cheese = cp.ingredientFactory.createCheese()
	cp.clam = cp.ingredientFactory.createClam()
}
