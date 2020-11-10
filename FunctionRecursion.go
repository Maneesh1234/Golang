package main

import "fmt"

type Employee struct {
	fname string
	lname string
}

func (emp Employee) fullname() {
	fmt.Println(emp.fname + " " + emp.lname)
}
func main() {
	e1 := Employee{"John", "Ponting"}
	e1.fullname()

}

// RETURN IN A FUNCTION
// func main() {
// 	x := add(4, 5)
// 	fmt.Println(x)
// }
func add(a int, b int) int {
	x := a + b
	return x
}

// PASSING VALUE IN A FUNCTION
// func main() {
// 	x := add(4, 5)
// 	fmt.Println(x)
// }
func add(a, b int) int {
	x := a + b
	return x
}

// PASSING VALUE IN A FUNCTION WHEN NUMBER OF ARGUMENT DOES NOT KNOW
// func main() {
// 	x := add(4, 5, 5)
// 	fmt.Println(x)
// }
func add(args ...int) int {
	s := 0
	for _, value := range args {
		s = s + value

	}

	return s
}

// RETURN MULTIPLE VALUE IN FUNCTION
// func main() {
// 	fmt.Println(addAll(10, 15, 20, 25, 30))
// }
func addAll(args ...int) (int, int) {
	finalAddValue := 0
	finalSubValue := 0
	for _, value := range args {
		finalAddValue += value
		finalSubValue -= value
	}
	return finalAddValue, finalSubValue
}

// RECURSION
// func main() {
// 	fmt.Println(factorial(5))
// }
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}

}
