package main

type MozzarellaCheese struct {
}

func NewMozzarellaCheese() Cheese {
	return &MozzarellaCheese{}
}

func (c *MozzarellaCheese) toString() string {
	return "Shredded Mozzarella"
}
