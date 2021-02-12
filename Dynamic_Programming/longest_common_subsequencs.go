// a file comparison program that outputs the differences between two files
// applications in bioinformatics
// LCS for input Sequences “ABCDGH” and “AEDFHR” is “ADH” of length 3.
// LCS for input Sequences “AGGTAB” and “GXTXAYB” is “GTAB” of length 4.
package main

import (
	"fmt"
)

// //USING RECURSION
// func main() {
// 	fmt.Println("Enter first string")
// 	var s1 string
// 	fmt.Scanln(&s1)
// 	fmt.Println("Enter Second string")
// 	var s2 string
// 	fmt.Scanln(&s2)
// 	fmt.Println(lcs(s1, s2, len(s1), len(s2)))

// }
// func lcs(str1 string, str2 string, i1 int, i2 int) int {
// 	if i1 == 0 || i2 == 0 {
// 		return 0
// 	}
// 	if str1[i1-1] == str2[i2-1] {
// 		return 1 + lcs(str1, str2, i1-1, i2-1)
// 	} else {
// 		return max(lcs(str1, str2, i1-1, i2), lcs(str1, str2, i1, i2-1))
// 	}
// }
// func max(num1 int, num2 int) int {
// 	if num1 > num2 {
// 		return num1
// 	}
// 	return num2
// }

//(2)USING DYNAMIC PROGRAMMING
func main() {
	fmt.Println("Enter first string")
	var s1 string
	fmt.Scanln(&s1)
	fmt.Println("Enter Second string")
	var s2 string
	fmt.Scanln(&s2)
	fmt.Println(lcs(s1, s2, len(s1), len(s2)))

}
func lcs(str1 string, str2 string, i1 int, i2 int) int {

	L := make([][]int, i1+1)
	for i := 0; i < i1+1; i++ {
		L[i] = make([]int, i2+1)
	}
	for i := i1 - 1; i >= 0; i-- {
		for j := i2 - 1; j >= 0; j-- {
			if str1[i] == str2[j] {
				L[i][j] = 1 + L[i+1][j+1]
			} else {
				L[i][j] = max(L[i+1][j], L[i][j+1])
			}
		}
	}

	return L[0][0]
}
func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
