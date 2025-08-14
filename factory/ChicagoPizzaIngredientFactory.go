package main

type ChicagoPizzaIngredientFactory struct {
}

func NewChicagoPizzaIngredientFactory() PizzaIngredientFactory {
	return &NYPizzaIngredientFactory{}
}

func (c *ChicagoPizzaIngredientFactory) createDough() Dough {
	return NewThickCrustDough()
}

func (c *ChicagoPizzaIngredientFactory) createSauce() Sauce {
	return NewPlumTomatoSauce()
}

func (c *ChicagoPizzaIngredientFactory) createCheese() Cheese {
	return NewMozzarellaCheese()
}

func (c *ChicagoPizzaIngredientFactory) createVeggies() []Veggies {
	veggies := []Veggies{
		NewBlackOlives(),
		NewSpinach(),
		NewEggplant(),
	}
	return veggies
}

func (c *ChicagoPizzaIngredientFactory) createPepperoni() Pepperoni {
	return NewSlicedPepperoni()
}

func (c *ChicagoPizzaIngredientFactory) createClam() Clams {
	return NewFrozenClams()
}
