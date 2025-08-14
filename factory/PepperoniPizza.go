package main

import "fmt"

type PepperoniPizza struct {
	ingredientFactory PizzaIngredientFactory

	PizzaSoul
}

func NewPepperoniPizza(ingredientFactory PizzaIngredientFactory) Pizza {
	pp := PepperoniPizza{ingredientFactory: ingredientFactory}
	return &pp
}

func (pp *PepperoniPizza) prepare() {
	fmt.Println("Preparing " + pp.name)
	pp.dough = pp.ingredientFactory.createDough()
	pp.sauce = pp.ingredientFactory.createSauce()
	pp.cheese = pp.ingredientFactory.createCheese()
	pp.veggies = pp.ingredientFactory.createVeggies()
	pp.pepperoni = pp.ingredientFactory.createPepperoni()
}
