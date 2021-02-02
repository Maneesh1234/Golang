package main

import (
	"fmt"
	//"test/counters"
	"gRpc/Oops/public_private"
)

func main() {
	// Create a variable of the exported type and
	// initialize the value to 10.
	// //(1) WE CAN ACCESS PUBLIC FIELD EASILY
	// //PUBLIC FIELD START WITH CAPITAL LETTERS
	// counter := public_private.AlertCounter(10)

	// fmt.Printf("Counter: %d\n", counter)

	// //(2) WE CAN NOT ACCESS PRIVATE FIELD EASIL
	// //PRIVATE FIELD START WITH SMALL LETTERS
	// //IT WILL GIVES ERROR
	// counter := public_private.alertCounter(10)

	// fmt.Printf("Counter: %d\n", counter)

	//(2) WE CAN ACCESS PRIVATE FIELD VIA PUBLIC METHOD IN ANOTHER FUNCTION
	counter := public_private.NewAlertCounter(10)

	fmt.Printf("Counter: %d\n", counter)
}
