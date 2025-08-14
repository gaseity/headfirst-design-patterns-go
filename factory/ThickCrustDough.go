package main

type ThickCrustDough struct {
}

func NewThickCrustDough() Dough {
	return &ThickCrustDough{}
}

func (d *ThickCrustDough) toString() string {
	return "ThickCrust style extra thick crust dough"
}
