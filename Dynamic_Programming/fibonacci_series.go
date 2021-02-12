package main

import (
	"fmt"
)

// //(1)	USING RECURSION
// func main() {
// 	fmt.Println("Enter a number")
// 	var num int
// 	fmt.Scanln(&num)
// 	for i := 0; i < num; i++ {
// 		fmt.Print(fibonacci(i), " ")
// 	}

// }
// func fibonacci(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	} else {
// 		return fibonacci(n-1) + fibonacci(n-2)
// 	}
// }

// //(2)	USING DYNAMIC (BUT IT IS USING EXTRA SPACE)
// func main() {

// 	fmt.Println("Enter a number")
// 	var num int
// 	fmt.Scanln(&num)

// 	fibonacci(num)

// }
// func fibonacci(n int) {
// 	if n == 1 {
// 		fmt.Print(0)
// 		return
// 	} else {
// 		a := []int{}
// 		for i := 0; i < n; i++ {
// 			if i == 0 || i == 1 {
// 				a = append(a, i)
// 			} else {
// 				a = append(a, a[i-1]+a[i-2])
// 			}
// 		}
// 		fmt.Print(a)
// 	}
// }

//(3)	USING DYNAMIC (WITHOUT USING EXTRA SPACE)
func main() {

	fmt.Println("Enter a number")
	var num int
	fmt.Scanln(&num)

	fibonacci(num)

}
func fibonacci(n int) {
	if n == 1 {
		fmt.Print(0)
		return
	} else {
		var next int
		prev := 0
		cur := 1
		for i := 0; i < n; i++ {
			if i == 0 {
				fmt.Print(prev, " ")
			} else {
				next = cur + prev
				prev = cur
				cur = next
				fmt.Print(prev, " ")
			}
		}

	}
}
