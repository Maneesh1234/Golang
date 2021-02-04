package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECTION := args[1]
	c, err := net.Dial("tcp", CONNECTION)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		//READ THE USER INPUT
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("CLIENT : ")
		text, _ := reader.ReadString('\n')
		//user input is sent to the TCP server over the network
		fmt.Fprintf(c, text+"\n")

		//read the TCP server’s response
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("SERVER SAYS: " + message)
		// terminate when you send the STOP command to the TCP server
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
