package main

import "fmt"

type VeggiePizza struct {
	ingredientFactory PizzaIngredientFactory

	PizzaSoul
}

func NewVeggiePizza(ingredientFactory PizzaIngredientFactory) Pizza {
	p := VeggiePizza{
		ingredientFactory: ingredientFactory,
	}
	return &p

}

func (vp *VeggiePizza) prepare() {
	fmt.Println("Preparing " + vp.name)
	vp.dough = vp.ingredientFactory.createDough()
	vp.sauce = vp.ingredientFactory.createSauce()
	vp.cheese = vp.ingredientFactory.createCheese()
	vp.veggies = vp.ingredientFactory.createVeggies()
}
