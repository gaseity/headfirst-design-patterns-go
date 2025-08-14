package main

import (
	"fmt"
	"strings"
)

type Pizza interface {
	prepare()

	bake()
	cut()
	box()

	setName(string)
	getName() string
	toString() string
}

type PizzaSoul struct {
	name string

	dough     Dough
	sauce     Sauce
	veggies   []Veggies
	cheese    Cheese
	pepperoni Pepperoni
	clam      Clams
}

func (ps *PizzaSoul) bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (ps *PizzaSoul) cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (ps *PizzaSoul) box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (ps *PizzaSoul) setName(name string) {
	ps.name = name
}

func (ps *PizzaSoul) getName() string {
	return ps.name
}

func (ps *PizzaSoul) toString() string {
	var result strings.Builder

	result.WriteString("---- " + ps.name + " ----\n")
	if ps.dough != nil {
		result.WriteString(ps.dough.toString())
		result.WriteString("\n")
	}
	if ps.sauce != nil {
		result.WriteString(ps.sauce.toString())
		result.WriteString("\n")
	}
	if ps.cheese != nil {
		result.WriteString(ps.cheese.toString())
		result.WriteString("\n")
	}
	if ps.veggies != nil {
		for _, veggie := range ps.veggies {
			result.WriteString(veggie.toString())
			result.WriteString(" ")
		}
		result.WriteString("\n")
	}
	if ps.clam != nil {
		result.WriteString(ps.clam.toString())
		result.WriteString("\n")
	}
	if ps.pepperoni != nil {
		result.WriteString(ps.pepperoni.toString())
		result.WriteString("\n")
	}
	return result.String()
}
