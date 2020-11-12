package main

import (
	"fmt"
)

// type person struct {
// 	firstName string
// 	lastName  string
// 	age       int
// }

// func main() {
// 	x := person{age: 22, firstName: "Maneesh", lastName: "Kumar"}
// 	fmt.Println(x)
// 	fmt.Println(x.firstName, "age is", x.age)
// }

// SIMPLE INHERITANCE USING STRUCT
type person struct {
	fname string
	lname string
}
type employee struct {
	person
	empId int
}

func (p person) details() {
	fmt.Println(p, " "+" I am a person")
}
func (e employee) details() {
	fmt.Println(e, " "+"I am a employee")
}
func main() {
	p1 := person{"Maneesh", "Kumar"}
	p1.details()
	e1 := employee{person: person{"Maneesh", "Kumar"}, empId: 14}
	e1.details()
}
