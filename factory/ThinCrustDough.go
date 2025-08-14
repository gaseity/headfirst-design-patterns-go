package main

type ThinCrustDough struct {
}

func NewThinCrustDough() Dough {
	return &ThinCrustDough{}
}

func (d *ThinCrustDough) toString() string {
	return "Thin Crust Dough"
}
