package main

type Command interface {
	execute()
	undo()
	getName() string
}
