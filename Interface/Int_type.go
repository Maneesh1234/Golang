// Interface Types: The interface is of two types one is static and another one is dynamic type. 
//The static type is the interface itself, for example, tank in the below example. But interface does 
//not have a static value so it always points to the dynamic values.
// Go program to illustrate the concept
// of dynamic values and types
package main

import "fmt"

// Creating an interface
type tank interface {

	// Methods
	Tarea() float64
	Volume() float64
}

func main() {

	var t tank
	fmt.Println("Value of the tank interface is: ", t)
	fmt.Printf("Type of the tank interface is: %T ", t)
}

