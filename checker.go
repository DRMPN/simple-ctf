package main

import (
	"os"
	"fmt"
	"net"
	"bufio"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: command host")
		os.Exit(1)
	}

	var COMMAND = os.Args[1]
	var HOST = os.Args[2]

	conn, _ := net.Dial("tcp", HOST)
	defer conn.Close()
	
	if message, _ := bufio.NewReader(conn).ReadString('\n'); message != "Enter your message: \n" {
		os.Exit(102)
	}
	
	if COMMAND == "check" {
		fmt.Fprintf(conn, "search\n")
		if message, _ := bufio.NewReader(conn).ReadString('\n'); message != "Enter pattern: \n" {
		os.Exit(102)
		}
		
		fmt.Println("OK")
	}

}