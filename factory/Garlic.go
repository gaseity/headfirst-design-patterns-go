package main

type Garlic struct {
}

func NewGarlic() Veggies {
	return &Garlic{}
}

func (g *Garlic) toString() string {
	return "Garlic"
}
