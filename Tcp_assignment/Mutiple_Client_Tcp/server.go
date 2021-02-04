package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

)



var count = 0


func handleConnection(conn net.Conn, count int) {
	
	for {

		netData, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		counter := strconv.Itoa(count)

		fmt.Print("CLIENT SAYS: ", counter, ": ", string(netData))

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("SERVER: ", counter, ": ")
		text, _ := reader.ReadString('\n')
		//user input is sent to the TCP server over the network
		fmt.Fprintf(conn, text+"\n")

	}
	
	conn.Close()
}
func main() {
	
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORTADDRESS := ":" + args[1]
	lis, err := net.Listen("tcp", PORTADDRESS)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
	
		count++
		go handleConnection(conn, count)

	
	}
}
