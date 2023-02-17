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

	if len(os.Args) < 3 {
		fmt.Println("Usage: checker.go host command key value")
		os.Exit(1)
	}

	var HOST = os.Args[1]
	var COMMAND = strings.ToLower(strings.TrimSpace(os.Args[2]))

	conn, err := net.Dial("tcp", HOST)
	if err!=nil {
		fmt.Printf("Cannot connect to specified host: %s\n", HOST)
		os.Exit(1)
	}
	defer conn.Close()

	var READER = bufio.NewReader(conn)
	
	if COMMAND == "check" {

		// TODO: argv

		if  message, _ := READER.ReadString('\n'); message != "Enter your message: \n" {
			fmt.Printf("No Enter your message: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "load\n")

		if message, _ := READER.ReadString('\n'); message != "Enter key: \n" {
			fmt.Printf("No Enter key: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "1\n")

		if  message, _ := READER.ReadString('\n'); message != "Value: 1\n" {
			fmt.Printf("No Value: %s\n", message)
			os.Exit(102)
		}

		if  message, _ := READER.ReadString('\n'); message != "Enter your message: \n" {
			fmt.Printf("No Enter your message: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "store\n")

		if message, _ := READER.ReadString('\n'); message != "Enter key: \n" {
			fmt.Printf("No Enter key: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "1\n")

		if message, _ := READER.ReadString('\n'); message != "Enter value: \n" {
			fmt.Printf("No Enter value: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "1\n")

		if message, _ := READER.ReadString('\n'); message != "Stored\n" {
			fmt.Printf("No Stored: %s\n", message)
			os.Exit(102)
		}

		if message, _ := READER.ReadString('\n'); message != "Enter your message: \n" {
			fmt.Printf("No Enter your message: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "search\n")

		if message, _ := READER.ReadString('\n'); message != "Enter pattern: \n" {
			fmt.Printf("No Enter pattern: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "1\n")

		if message, _ := READER.ReadString('\n'); message != "Found: 111, 11, 1\n" {
			fmt.Printf("No Found: %s\n", message)
			os.Exit(102)
		}

		if message, _ := READER.ReadString('\n'); message != "Enter your message: \n" {
			fmt.Printf("No Enter your message: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "WRONG\n")

		if message, _ := READER.ReadString('\n'); message != "Unkown command: wrong\n" {
			fmt.Printf("No Unkown command:: %s\n", message)
			os.Exit(102)
		}
		log.Printf("%s: ok", COMMAND)
	} else if COMMAND == "put" { 
		fmt.Println("put")
		// go run checker.go localhost:7777 put 1 1 
		key := os.Args[3]
		value := os.Args[4]
		log.Printf("%s: %s - %s", COMMAND, key, value)
		// Enter your message: 
		// store
		// Enter key: 
		// 1
		// Enter value: 
		// 1
		// Stored
	} else if COMMAND == "get" { 
		
		// TODO: argv

		key := os.Args[3]+"\n"
		log.Printf("%s: %s", COMMAND, key)

		if message, _ := READER.ReadString('\n'); message != "Enter your message: \n" {
			fmt.Printf("No Enter your message: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "load\n")

		if message, _ := READER.ReadString('\n'); message != "Enter key: \n" {
			fmt.Printf("No Enter key: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, key)

		if message, _ := READER.ReadString('\n'); message == "No value\n" {
			fmt.Printf("No flag is found: %s\n", message)
			os.Exit(102)
		} 

		if message, _ := READER.ReadString('\n'); message != "Enter your message: \n" {
			fmt.Printf("No Enter your message: %s\n", message)
			os.Exit(102)
		}

		log.Printf("%s: ok", COMMAND)
	} else {
		fmt.Println("Usage: command host")
		os.Exit(1)
	}

	os.Exit(0)
}