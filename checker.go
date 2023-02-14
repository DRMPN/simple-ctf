package main

import (
	"os"
	"log"
	"fmt"
	"net"
	"bufio"
	"strings"
)

func main() {

	// usage: nc [-46CDdFhklNnrStUuvZz] [-I length] [-i interval] [-M ttl]
	// 		[-m minttl] [-O length] [-P proxy_username] [-p source_port]
	// 		[-q seconds] [-s source] [-T keyword] [-V rtable] [-W recvlimit] [-w timeout]
	// 		[-X proxy_protocol] [-x proxy_address[:port]]           [destination] [port]

	if len(os.Args) != 3 {
		fmt.Println("Usage: command host")
		os.Exit(1)
	}

	var COMMAND = strings.ToLower(strings.TrimSpace(os.Args[1]))
	var HOST = os.Args[2]

	conn, err := net.Dial("tcp", HOST)
	if err!=nil {
		fmt.Printf("Cannot connect to specified host: %s\n", HOST)
		os.Exit(1)
	}
	defer conn.Close()

	var READER = bufio.NewReader(conn)
	
	if COMMAND == "check" {
		if  message, _ := READER.ReadString('\n'); message != "Enter your message: \n" {
			log.Printf("Error 13: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "load\n")

		if message, _ := READER.ReadString('\n'); message != "Enter key: \n" {
			log.Printf("Error 5: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "1\n")

		if  message, _ := READER.ReadString('\n'); message != "Value: 1\n" {
			log.Printf("Error 11: %s\n", message)
			os.Exit(102)
		}

		if  message, _ := READER.ReadString('\n'); message != "Enter your message: \n" {
			log.Printf("Error 10: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "store\n")

		if message, _ := READER.ReadString('\n'); message != "Enter key: \n" {
			log.Printf("Error 3: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "1\n")

		if message, _ := READER.ReadString('\n'); message != "Enter value: \n" {
			log.Printf("Error 0: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "1\n")

		if message, _ := READER.ReadString('\n'); message != "Stored\n" {
			log.Printf("Error 1: %s\n", message)
			os.Exit(102)
		}

		if message, _ := READER.ReadString('\n'); message != "Enter your message: \n" {
			log.Printf("Error 14: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "search\n")

		if message, _ := READER.ReadString('\n'); message != "Enter pattern: \n" {
			log.Printf("Error 2: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "1\n")

		if message, _ := READER.ReadString('\n'); message != "Found: 111, 11, 1\n" {
			log.Printf("Error 18: %s\n", message)
			os.Exit(102)
		}

		if message, _ := READER.ReadString('\n'); message != "Enter your message: \n" {
			log.Printf("Error 20: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "WRONG\n")

		if message, _ := READER.ReadString('\n'); message != "Unkown command: wrong\n" {
			log.Printf("Error 20: %s\n", message)
			os.Exit(102)
		}
		log.Printf("%s: ok", COMMAND)
	} else {
		fmt.Println("Usage: command host")
		os.Exit(1)
	}
	//comments
	//get
	//put
	os.Exit(0)
}