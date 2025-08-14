package main

type SlicedPepperoni struct {
}

func NewSlicedPepperoni() Pepperoni {
	return &SlicedPepperoni{}
}

func (p *SlicedPepperoni) toString() string {
	return "Sliced Pepperoni"
}
