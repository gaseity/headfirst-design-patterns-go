package main

type ReggianoCheese struct {
}

func NewReggianoCheese() Cheese {
	return &ReggianoCheese{}
}

func (c *ReggianoCheese) toString() string {
	return "Reggiano Cheese"
}
