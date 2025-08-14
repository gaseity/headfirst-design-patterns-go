package main

import "fmt"

func main() {
	nyStore := NewNYPizzaStore()
	chicagoStore := NewChicagoPizzaStore()

	pizza := nyStore.orderPizza("cheese")
	fmt.Println("Ethan ordered a " + pizza.toString() + "\n")

	pizza = chicagoStore.orderPizza("cheese")
	fmt.Println("Joel ordered a " + pizza.toString() + "\n")

	pizza = nyStore.orderPizza("clam")
	fmt.Println("Ethan ordered a " + pizza.toString() + "\n")

	pizza = chicagoStore.orderPizza("clam")
	fmt.Println("Joel ordered a " + pizza.toString() + "\n")

	pizza = nyStore.orderPizza("pepperoni")
	fmt.Println("Ethan ordered a " + pizza.toString() + "\n")

	pizza = chicagoStore.orderPizza("pepperoni")
	fmt.Println("Joel ordered a " + pizza.toString() + "\n")

	pizza = nyStore.orderPizza("veggie")
	fmt.Println("Ethan ordered a " + pizza.toString() + "\n")

	pizza = chicagoStore.orderPizza("veggie")
	fmt.Println("Joel ordered a " + pizza.toString() + "\n")
}
