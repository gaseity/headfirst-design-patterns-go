package main

type FreshClams struct {
}

func NewFreshClams() Clams {
	return &FreshClams{}
}

func (c *FreshClams) toString() string {
	return "Fresh Clams from Long Island Sound"
}
