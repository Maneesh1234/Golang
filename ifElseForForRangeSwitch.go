var a int = 12
	/* check the boolean condition using if statement */
	if a%2 == 0 { /* if condition is true then print the following
		 */fmt.Printf("a is even number")
	}

	var a int = 15
	/* check the boolean condition */
	if a%2 == 0 {
		/* if condition is true then print the following */
		fmt.Printf("a is even\n")
	} else {
		/* if condition is false then print the following */
		fmt.Printf("a is odd\n")
	}
	fmt.Printf("value of a is : %d\n", a)

	fmt.Print("Enter number: ")
	var input int
	fmt.Scanln(&input)
	fmt.Print(input)
	/* check the boolean condition */
	if input%2 == 0 {
		/* if condition is true then print the following */
		fmt.Printf(" is even\n")
	} else {
		/* if condition is false then print the following */
		fmt.Printf(" is odd\n")
	}

	fmt.Print("Enter text: ")
	var i int
	fmt.Scanln(&i)
	if i < 0 || i > 100 {
		fmt.Print("Please enter valid no")
	} else if i >= 0 && i < 50 {
		fmt.Print(" Fail")
	} else if i >= 50 && i < 60 {
		fmt.Print(" D Grade")
	} else if i >= 60 && i < 70 {
		fmt.Print(" C Grade")
	} else if i >= 70 && i < 80 {
		fmt.Print(" B Grade")
	} else if i >= 80 && i < 90 {
		fmt.Print(" A Grade")
	} else if i >= 90 && i <= 100 {
		fmt.Print(" A+ Grade")
	}

	fmt.Println("Enter Your First Name: ")

	// var then variable name then variable type
	var first string

	// Taking input from user
	fmt.Scanln(&first)
	fmt.Println("Enter Second Last Name: ")
	var second string
	fmt.Scanln(&second)

	Print function is used to
	display output in the same line
	fmt.Print("Your Full Name is: ")

	// Addition of two string
	fmt.Print(first + " " + second)

	var x int = 10
	var y int = 9
	/* check the boolean condition */
	if x >= 10 {
		/* if condition is true then check the following */
		if y >= 10 {
			/* if condition is true then print the following */
			fmt.Printf("Inside nested If Statement \n")
		}
		fmt.Printf("Inside If Statement \n")
	}
	fmt.Printf("Value of x is : %d\n", x)
	fmt.Printf("Value of y is : %d\n", y)

	fmt.Print("Enter Number: ")
	var input int
	fmt.Scanln(&input)
	switch input {
	case 10, 15:
		fmt.Print("the value is 10 or 15")
	case 20:
		fmt.Print("the value is 20")
	case 30:
		fmt.Print("the value is 30")
	case 40:
		fmt.Print("the value is 40")
	default:
		fmt.Print(" It is not 10,15 ,20,30,40 ")
	}

	k := 20
	switch k {
	case 10:
		fmt.Println("was <= 10")
		fallthrough
	case 20:
		fmt.Println("was <= 20")
		fallthrough
	case 30:
		fmt.Println("was <= 30")
		fallthrough
	case 40:
		fmt.Println("was <= 40")
		fallthrough
	case 50:
		fmt.Println("was <= 50")
		fallthrough
	case 60:
		fmt.Println("was <= 60")
		fallthrough
	default:
		fmt.Println("default case")
	}

	COUNTER FOR LOOP
	for a := 0; a < 10; a++ {
		fmt.Print(a, "\n")
	}

	//NESTED FOR LOOP
	for a := 0; a < 3; a++ {
		for b := 3; b > 0; b-- {
			fmt.Print(a, " ", b, "\n")
		}
	}

	//INFINITE FOR LOOP
	for true {
		fmt.Printf("This loop will run forever.\n")
	}

	//CONDITIONER FOR LOOP IT IS SIMILAR TO WHILE LOOP OF OTHER LANGUAGE
	sum := 1
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}

	//FOR RANGE LOOP . IT IS SIMILAR FOR EACH LOOP
	nums := []int{2, 3, 4}
	sum := 0
	for _, value := range nums { // "_ " is to ignore the index
		sum += value
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	// map data structure
	kvs := map[string]string{"1": "mango", "2": "apple", "3": "banana"}
	for j, v := range kvs {
		fmt.Printf("%s -> %s\n", j, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
	for i, c := range "Hi" {
		fmt.Println(i, c)
	}
