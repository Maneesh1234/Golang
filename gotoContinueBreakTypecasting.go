package main

func main() {
	// goto  KEYWORD
		var input int
	Loop:
		fmt.Print("Enter your age ")
		fmt.Scanln(&input)
		if input <= 18 {
			fmt.Print("You are not eligible to vote ")
			goto Loop
		} else {
			fmt.Print("You can vote ")
		}

	// break KEYWORD
	var a int = 1
	for a < 10 {
		fmt.Print("Value of a is ", a, "\n")
		a++
		if a > 5 {
			/* terminate the loop using break statement */
			fmt.Print("Only less than 5  value left \n")
			break
		}
	}

	var a int
	var b int
	for a = 1; a <= 3; a++ {
		for b = 1; b <= 3; b++ {
			if a == 1 && b == 1 {
				break
			}
			fmt.Print(a, " ", b, "\n")
		}
	}

	// continue KEYWORD
	var a int = 1
	for a < 10 {
		if a == 5 {
			/* skip the iteration */
			a = a + 1
			continue
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}

	var a int = 1
	var b int = 1
	/* do loop execution */
	for a = 1; a < 3; a++ {
		for b = 1; b < 3; b++ {
			if a == 2 && b == 2 {
				/* skip the iteration */
				continue
			}
			fmt.Printf("value of a and b is %d %d\n", a, b)
		}
		fmt.Printf("value of a and b is %d %d\n", a, b)
	}

	//const KEYWORD
	const HEIGHT int = 100
	const WIDTH int = 200
	var area int
	area = HEIGHT * WIDTH
	fmt.Printf("value of area : %d", area)

  // CONVERSION OF DATA TYPE
	var i int = 102
	var f float64 = 7.56
	var s1 string = "156"
	var s2 string = "89.58"

	fmt.Println(float64(i))
	fmt.Println(int(f))

	newInt, _ := strconv.ParseInt(s1, 0, 64)
	fmt.Println(newInt)

	newfloat, _ := strconv.ParseFloat(s2, 64)
	fmt.Println(newfloat)
  }
