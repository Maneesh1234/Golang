
package main

// import (
// 	"fmt"
// )

import (
	"fmt"
)

var(
	actorName string="Elisabeth Sladen "
	companion string="Sarah John Smith"
	docternumber int =3
	season int =11

)

// var{

// }

lower case letter variable compiler expose the variable at package level
var i int = 81005

upper case letter variable compiler expose the variable to outside world
var I int =58

var i float32 = "85"      //it gives error

func main() {
	fmt.Println("Hello World how!")

	fmt.Println(42)

	var i int
	i = 42
	i = 27

	//variable visible at block level
	var i int = 85

	i := 75

	fmt.Println(i)

	var i int = 85

	var i float32 = 85

	i := 75    // it is declare and initialize the variable

	i := 75.

	var i int = 85

	i = 75

	var theURL string = "http://google.com"
	fmt.Printf("%v, %T", theURL, theURL)

	var i int = 42
	fmt.Printf("%v, %T\n", i, i)
	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var i float32 = 42.5
	fmt.Printf("%v, %T\n", i, i)
	var j int
	j = int(i)
	fmt.Printf("%v, %T\n", j, j)

	//it gives error because go does not want to loose information
	var i float32 = 42.5
	fmt.Printf("%v, %T\n", i, i)
	var j int
	j = i
	fmt.Printf("%v, %T\n", j, j)

	// It print unicode value of i
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)
	var j string
	j = string(i)
	fmt.Printf("%v, %T\n", j, j)

	// It convert integer to string
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)
	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)

	// BOOLEAN TYPE
	var i bool = true
	fmt.Printf("%v, %T\n", i, i)

	i := 1 == 1
	j := 1 == 2
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)

	// Every time you intialize variable default value it assign 0 value
	var i bool
	fmt.Printf("%v, %T\n", i, i)

	NUMERIC TYPE
	UNSIGNED INTEGER
	var i uint16 = 42
	fmt.Printf("%v, %T\n", i, i)

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(a * b)

	// It gives error because mismatched type
	var a int = 10
	var b int8 = 3
	fmt.Println(a + b)

	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b))

	// BITWISE OPERATOR
	a := 10            // 1010
	b := 3             //0011
	fmt.Println(a & b) // 0010
	fmt.Println(a | b) //1011
	fmt.Println(a ^ b)     // 1001
	fmt.Println(a &^ b)     //1000

	b := 8
	fmt.Println(b << 3)
	fmt.Println(b >> 3)

	b := 5.12
	b = 2.12e12
	b = 4.12e12
	fmt.Printf("%v, %T\n", b, b)

	a := 10.5
	b := 3.1
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)

	// COMPLEX NUMBER
	var i complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", i, i)

	var i complex64 = 2i
	fmt.Printf("%v, %T\n", i, i)

	a := 1 + 2i
	b := 2 + 5.2i
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)

	var i complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", real(i), real(i))
	fmt.Printf("%v, %T\n", imag(i), imag(i))

	var i complex128 = complex(2, 4)
	fmt.Printf("%v, %T\n", i, i)

	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)

	s := "this is a string"
	fmt.Printf("%v, %T\n", s[2], s[2])

	s := "this is a string"
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2]))

	// STRING IS UNMUTABLE
	s := "this is a string"
	s[2] = "u"
	fmt.Printf("%v, %T\n", s, s)

	s := "this is a string"
	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	s := "this is a string"
	// It convert string to ascii character of each character
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%T %T %T %T\n", i, f, b, s)              // Prints type of the variable
	fmt.Printf("%v   %v      %v  %q     \n", i, f, b, s) //prints initial value of the variable

}
