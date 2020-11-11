package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "Good Morning"
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}

// len FUNCTION IN STRING
func main() {
	str := "I love my country"
	fmt.Println(len(str))
}

// ASCII VALUE A
func main() {
	fmt.Println("Ascii value of A is ", "A"[0])
}

// VARIOUS FUNCTION IN STRING
func main() {
	str := "India"
	// CONVERT INTO UPPER CASE
	fmt.Println(strings.ToUpper(str))
	// CONVERT INTO LOWER CASE
	fmt.Println(strings.ToLower(str))
	// CHECK PREFIX OF THE STRING
	fmt.Println(strings.HasPrefix(str, "In"))
	// CHECK SUFFIX OF THE STRING
	fmt.Println(strings.HasSuffix(str, "ia"))
	var arr = []string{"a", "b", "c", "d"}
	fmt.Println(strings.Join(arr, "*"))
	fmt.Println(strings.Repeat(str, 4))
	fmt.Println(strings.Contains(str, "nd"))
	fmt.Println(strings.Index(str, "nd"))
	fmt.Println(strings.Count(str, "i"))
	fmt.Println(strings.Replace(str, "i", "Z", 2))
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
	fmt.Println(strings.TrimSpace(" \t\n I love my country  \n\t\r\n"))
	fmt.Println(strings.ContainsAny("Hello", "A"))
	fmt.Println(strings.ContainsAny("Hello", "o & e"))
	fmt.Println(strings.ContainsAny("Hello", ""))
	fmt.Println(strings.ContainsAny("", ""))
}

func main() {
	str := "I,love,my,country"
	var arr []string = strings.Split(str, ",")
	fmt.Println(len(arr))
	for i, v := range arr {
		fmt.Println("Index : ", i, "value : ", v)
	}
}
