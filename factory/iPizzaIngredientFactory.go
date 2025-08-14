package main

type PizzaIngredientFactory interface {
	createDough() Dough
	createSauce() Sauce
	createCheese() Cheese
	createVeggies() []Veggies
	createPepperoni() Pepperoni
	createClam() Clams
}
