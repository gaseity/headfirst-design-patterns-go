package main

type ChicagoPizzaStore struct {
	ingredientFactory PizzaIngredientFactory

	PizzaStoreSoul
}

func NewChicagoPizzaStore() PizzaStore {
	return &ChicagoPizzaStore{
		ingredientFactory: NewChicagoPizzaIngredientFactory(),
	}
}

func (cs *ChicagoPizzaStore) createPizza(item string) Pizza {
	var pizza Pizza

	if item == "cheese" {
		pizza = NewCheesePizza(cs.ingredientFactory)
		pizza.setName("Chicago Style Cheese Pizza")
	} else if item == "veggie" {
		pizza = NewVeggiePizza(cs.ingredientFactory)
		pizza.setName("Chicago Style Veggie Pizza")
	} else if item == "clam" {
		pizza = NewClamPizza(cs.ingredientFactory)
		pizza.setName("Chicago Style Clam Pizza")
	} else if item == "pepperoni" {
		pizza = NewPepperoniPizza(cs.ingredientFactory)
		pizza.setName("Chicago Style Pepperoni Pizza")
	}
	return pizza
}

func (cs *ChicagoPizzaStore) orderPizza(item string) Pizza {
	pizza := cs.createPizza((item))
	cs.cookPizza(pizza)
	return pizza
}
