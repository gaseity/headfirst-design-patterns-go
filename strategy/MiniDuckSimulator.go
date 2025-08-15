package main

import "fmt"

var ducks []string = []string{
	"Mallard",
	"Model",
}

func main() {
	for _, duck := range ducks {

		d := NewDuck(duck)
		if d == nil {
			fmt.Println("Duck cooked ", duck)
			continue
		}
		fmt.Println("----")

		d.Display()
		d.Fly()
		d.Quack()
		d.Swim()

	}

}
