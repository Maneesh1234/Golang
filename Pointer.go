package main

import (
	"fmt"
)

func main() {
	x := 6

	fmt.Println("Before change x", x)
	changeX(&x)
	fmt.Println("After change x ", x)
}
func changeX(x *int) {
	*x = 85
}

// func main() {
// 	ptr := new(int)
// 	fmt.Println("Before change ptr", *ptr)
// 	changePtr(ptr)
// 	fmt.Println("After change ptr", *ptr)
// }
// func changePtr(ptr *int) {
// 	*ptr = 10
// }
