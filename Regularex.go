package main

import (
	"fmt"
	"regexp"
)

// func main() {
// 	re := regexp.MustCompile(".com")
// 	fmt.Println(re.FindString("facebook.com"))
// 	fmt.Println(re.FindString("geekforgeeks.org"))
// 	fmt.Println(re.FindString("xy.com"))
// }

// func main() {
// 	re := regexp.MustCompile(".com")
// 	fmt.Println(re.FindStringIndex("facebook.com"))
// 	fmt.Println(re.FindStringIndex("geekforgeeks.org"))
// 	fmt.Println(re.FindStringIndex("xy.com"))
// }

func main() {
	re := regexp.MustCompile("f([a-z]+)ing")
	fmt.Println(re.FindStringSubmatch("flying"))
	fmt.Println(re.FindStringSubmatch("abcfloatingxyz"))
}
