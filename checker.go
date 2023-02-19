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

	if len(os.Args) < 3 {
		fmt.Println("Usage: checker [HOST] ([CHECK], [PUT KEY VALUE]), [GET KEY VALUE])")
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

		if len(os.Args) != 3 {
			fmt.Println("Usage: checker.go [HOST] [CHECK]")
			os.Exit(1)
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
		
		if  message, _ := READER.ReadString('\n'); message != "Enter your message: \n" {
			fmt.Printf("No Enter your message: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "load\n")

		if message, _ := READER.ReadString('\n'); message != "Enter key: \n" {
			fmt.Printf("No Enter key`: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "1\n")

		if  message, _ := READER.ReadString('\n'); message != "Value: 1\n" {
			fmt.Printf("No Value: %s\n", message)
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

		if message, _ := READER.ReadString('\n'); message != "Found: 1\n" {
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

		if len(os.Args) != 5 {
			fmt.Println("Usage: checker.go [HOST] [PUT] [KEY] [VALUE]")
			os.Exit(1)
		}

		key := os.Args[3]+"\n"
		value := os.Args[4]+"\n"

		if message, _ := READER.ReadString('\n'); message != "Enter your message: \n" {
			fmt.Printf("No Enter your message: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, "store\n")

		if message, _ := READER.ReadString('\n'); message != "Enter key: \n" {
			fmt.Printf("No Enter key: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, key)

		if message, _ := READER.ReadString('\n'); message != "Enter value: \n" {
			fmt.Printf("No Enter key: %s\n", message)
			os.Exit(102)
		}
		fmt.Fprintf(conn, value)

		if message, _ := READER.ReadString('\n'); message != "Stored\n" {
			fmt.Printf("No store: %s\n", message)
			os.Exit(102)
		} 

		log.Printf("%s: ok", COMMAND)
	} else if COMMAND == "get" { 
		
		if len(os.Args) != 4 {
			fmt.Println("Usage: checker.go [HOST] [GET] [VALUE]")
			os.Exit(1)
		}

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
	os.Exit(101)
}