package main

type Observer interface {
	Update(temp, humidity, pressure int)
}
