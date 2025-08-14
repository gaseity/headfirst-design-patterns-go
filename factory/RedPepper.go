package main

type RedPepper struct {
}

func NewRedPepper() Veggies {
	return &RedPepper{}
}

func (m *RedPepper) toString() string {
	return "Red Pepper"
}
