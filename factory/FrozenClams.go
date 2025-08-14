package main

type FrozenClams struct {
}

func NewFrozenClams() Clams {
	return &FrozenClams{}
}

func (c *FrozenClams) toString() string {
	return "Frozen Clams from Chesapeake Bay"
}
