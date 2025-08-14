package main

type Onion struct {
}

func NewOnion() Veggies {
	return &Onion{}
}

func (o *Onion) toString() string {
	return "Onion"
}
