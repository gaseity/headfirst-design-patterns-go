package main

type NYPizzaIngredientFactory struct {
}

func NewNYPizzaIngredientFactory() PizzaIngredientFactory {
	return &NYPizzaIngredientFactory{}
}

func (n *NYPizzaIngredientFactory) createDough() Dough {
	return NewThinCrustDough()
}

func (n *NYPizzaIngredientFactory) createSauce() Sauce {
	return NewMarinaraSauce()
}

func (n *NYPizzaIngredientFactory) createCheese() Cheese {
	return NewReggianoCheese()
}

func (n *NYPizzaIngredientFactory) createVeggies() []Veggies {
	veggies := []Veggies{
		NewGarlic(),
		NewOnion(),
		NewMushroom(),
		NewRedPepper(),
	}
	return veggies
}

func (n *NYPizzaIngredientFactory) createPepperoni() Pepperoni {
	return NewSlicedPepperoni()
}

func (n *NYPizzaIngredientFactory) createClam() Clams {
	return NewFreshClams()
}
