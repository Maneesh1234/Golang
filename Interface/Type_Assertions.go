// Type Assertions: In Go language, type assertion is an operation 
//applied to the value of the interface. Or in other words, type assertion is a process to extract the values of the interface.
// Go program to illustrate
// the type assertion
package main

import "fmt"

func myfun(a interface{}) {

	// Extracting the value of a
	val := a.(string)
	fmt.Println("Value: ", val)
}
func main() {

	var val interface {
	} = "Maneesh Kumar"
	
	myfun(val)
}

//  Go program to illustrate type assertion
package main

import "fmt"

func myfun(a interface{}) {
	value, ok := a.(float64)
	fmt.Println(value, ok)
}
func main() {

	var a1 interface {
	} = 98.09

	myfun(a1)

	var a2 interface {
	} = "GeeksforGeeks"

	myfun(a2)
}
