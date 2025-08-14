package main

type BlackOlives struct {
}

func NewBlackOlives() Veggies {
	return &BlackOlives{}
}

func (m *BlackOlives) toString() string {
	return "Black Olives"
}
