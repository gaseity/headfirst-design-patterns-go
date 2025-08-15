package main

func main() {
	boiler := GetInstance()
	boiler.fill()
	boiler.boil()
	boiler.drain()

	// will return the existing instance
	boiler2 := GetInstance()
	boiler2.fill()
	boiler2.boil()
	boiler2.drain()
}
