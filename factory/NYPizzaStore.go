package main

type NYPizzaStore struct {
	ingredientFactory PizzaIngredientFactory

	PizzaStoreSoul
}

func NewNYPizzaStore() PizzaStore {
	return &NYPizzaStore{
		ingredientFactory: NewNYPizzaIngredientFactory(),
	}
}

func (nys *NYPizzaStore) createPizza(item string) Pizza {

	var pizza Pizza

	if item == "cheese" {
		pizza = NewCheesePizza(nys.ingredientFactory)
		pizza.setName("New York Style Cheese Pizza")
	} else if item == "veggie" {
		pizza = NewVeggiePizza(nys.ingredientFactory)
		pizza.setName("New York Style Veggie Pizza")
	} else if item == "clam" {
		pizza = NewClamPizza(nys.ingredientFactory)
		pizza.setName("New York Style Clam Pizza")
	} else if item == "pepperoni" {
		pizza = NewPepperoniPizza(nys.ingredientFactory)
		pizza.setName("New York Style Pepperoni Pizza")
	}
	return pizza
}

func (nys *NYPizzaStore) orderPizza(item string) Pizza {
	pizza := nys.createPizza(item)
	nys.makingPizza(pizza)
	return pizza
}
