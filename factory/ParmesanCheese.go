package main

type ParmesanCheese struct {
}

func NewParmesanCheese() Cheese {
	return &ParmesanCheese{}
}

func (c *ParmesanCheese) toString() string {
	return "Shredded Parmesan"
}
