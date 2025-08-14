package main

type Mushroom struct {
}

func NewMushroom() Veggies {
	return &Mushroom{}
}

func (m *Mushroom) toString() string {
	return "Mushrooms"
}
