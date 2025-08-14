package main

type PlumTomatoSauce struct {
}

func NewPlumTomatoSauce() Sauce {
	return &PlumTomatoSauce{}
}

func (m *PlumTomatoSauce) toString() string {
	return "Tomato sauce with plum tomatoes"
}
