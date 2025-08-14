package main

type Eggplant struct {
}

func NewEggplant() Veggies {
	return &Eggplant{}
}

func (m *Eggplant) toString() string {
	return "Eggplant"
}
