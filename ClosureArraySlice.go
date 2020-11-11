package main

func main() {
	number := 10
	addNum := func() int {
		number += number
		return number
	}
	fmt.Println(addNum())
	fmt.Println(addNum())
}
ARRAY
func main() {
	var x [5]int
	var i, j int
	for i = 0; i < 5; i++ {
		x[i] = i + 10
	}
	for j = 0; j < 5; j++ {
		fmt.Printf("Element[%d] = %d\n", j, x[j])
	}
}

// MULTIDIMENSIONAL ARRAY
func main() {
	/* an array with 3 rows and 3 columns*/
	var a = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var i, j int
	/* output each array element's value */
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}

// SLICE DATA TYPE
func main() {
	odd := [6]int{2, 4, 6, 8, 10, 12}
	var s []int = odd[1:4]
	fmt.Println(s)
}

// SLICE IS THE REFERENCE WHEN CHANGE IN SLICE THEN IT ALSO AFFECT IN ARRAY
func main() {
	names := [4]string{
		"John",
		"Jim",
		"Jack",
		"jen",
	}
	fmt.Println(names)
	slice1 := names[0:2]
	slice2 := names[1:3]
	fmt.Println(slice1, slice2)
	slice2[0] = "ZZZ"
	fmt.Println(slice1, slice2)
	fmt.Println(names)
}

// Slice Literal
func main() {
	s := []struct {
		i int
		b bool
		x int
	}{
		{1, true, 9},
		{2, false, 8},
		{3, true, 7},
		{4, true, 4},
		{5, false, 3},
		{6, true, 2},
	}
	fmt.Println(s)
}

// LOWER BOUND AND UPPER BOUNDS IN SLICES
func main() {
	slice1 := []int{10, 5, 4, 8, 9, 7, 5, 6, 8}
	slice2 := slice1[2:4]
	fmt.Println(slice2)
	slice3 := slice1[:3]
	fmt.Println(slice3)
	slice4 := slice1[2:]
	fmt.Println(slice4)
	fmt.Println(slice1)
}

// CALCULATING LENGTH AND CAPACITY OF SLICE
func main() {
	slice1 := []int{2, 4, 6, 8, 10, 12, 14}
	printSlice(slice1)
	// Slice the slice to give it zero length.
	slice2 := slice1[:0]
	printSlice(slice2)
	// Extend its length.
	slice3 := slice1[:4]
	printSlice(slice3)
	// Drop its first two values.
	slice4 := slice1[2:]
	printSlice(slice4)
}
func printSlice(s []int) {
	fmt.Printf("length =%d capacity=%d %v\n", len(s), cap(s), s)
}

//  SLICE MAKE FUNCTION
func main() {
	slice := make([]int, 10)
	printSlice("slice", slice)
	slice1 := make([]int, 0, 10)
	printSlice("slice1", slice1)
	slice2 := slice1[:5]
	printSlice("slice2", slice2)
	slice3 := slice2[2:5]
	printSlice("slice3", slice3)
 }
 func printSlice(s string, x []int) {
	fmt.Printf("%s length=%d capacity=%d %v\n",
	   s, len(x), cap(x), x)
 }

// Command Line Argument
func main() {
	var s string
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + " "
	}
	fmt.Println(s)
}

func main() {
	arumentWithPath := os.Args  //returns all arguments including path
	arumentSlice := os.Args[1:] //returns all elements after path
	arumentAt2 := os.Args[2]    //returns specified argument only
	fmt.Println(arumentWithPath)
	fmt.Println(arumentSlice)
	fmt.Println(arumentAt2)
}
