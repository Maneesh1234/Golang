// // (1)Golang program to illustrate the
// // concept of interfaces
// package main

// import (
// 	"fmt"
// )

// // defining an interface
// type Sport interface {

// 	// name of sport method
// 	sportName() string
// }

// // declaring a struct
// type Human struct {

// 	// defining struct variables
// 	name  string
// 	sport string
// }

// // function to print book details
// func (h Human) sportName() string {

// 	// returning a string value
// 	return h.name + " plays " + h.sport + "."
// }

// // main function
// func main() {

// 	// declaring a struct instance
// 	human1 := Human{"Rahul", "chess"}

// 	// printing details of human1
// 	fmt.Println(human1.sportName())

// 	// declaring another struct instance
// 	human2 := Human{"Riya", "carrom"}

// 	// printing details of human2
// 	fmt.Println(human2.sportName())
// }

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func TotalPerimeter(shapes ...Shape) float64 {
	var peri float64
	for _, s := range shapes {
		peri += s.Perimeter()
	}
	return peri
}

func main() {
	r := Rect{width: 10, height: 10}
	c := Circle{radius: 10}
	fmt.Println("Total Area: ", TotalArea(r, c))
	fmt.Println("Total Perimeter: ", TotalPerimeter(r, c))
}
