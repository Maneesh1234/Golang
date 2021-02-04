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
	q := []string{}
	for {

		//READ THE USER INPUT
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("CLIENT  1: ")
		text, _ := reader.ReadString('\n')
		q = append(q, text)
		//user input is sent to the TCP server over the network
		fmt.Fprintf(c, text+"\n")

		//read the TCP server’s response
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("SERVER SAYS: " + message + "\n")
		q = append(q, message)
		fmt.Println("queue is :", q)
		q = q[1:]
		fmt.Println("updated queue is :", q)
		//q=q[1:]
		// terminate when you send the STOP command to the TCP server
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
