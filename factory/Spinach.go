package main

type Spinach struct {
}

func NewSpinach() Veggies {
	return &Spinach{}
}

func (s *Spinach) toString() string {
	return "Spinach"
}
