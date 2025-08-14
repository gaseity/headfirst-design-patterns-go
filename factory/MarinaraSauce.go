package main

type MarinaraSauce struct {
}

func NewMarinaraSauce() Sauce {
	return &MarinaraSauce{}
}

func (m *MarinaraSauce) toString() string {
	return "Marinara Sauce"
}
