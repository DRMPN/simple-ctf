package main

import (
	"os"
	"log"
	"fmt"
	"net"
	"bufio"
	"strings"
	"io/ioutil"
)

var DIRNAME = "DB"
var PORT = ":7777"

func main() {
	if fileinfo, err := os.Stat(DIRNAME); err!=nil || !fileinfo.IsDir(){
		log.Printf("Created directory: %s\n", DIRNAME)
		os.Mkdir(DIRNAME, os.ModePerm)
	}

	ln, _ := net.Listen("tcp", PORT)
	defer ln.Close()

	log.Printf("Service started on port %s", PORT)
	for {
		log.Println("Waiting for connection")
		conn, _ := ln.Accept()
		defer conn.Close()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer log.Println("Connection closed")
	log.Println("Connection established")
	for{
		log.Println("Waiting for user message")
		conn.Write([]byte("Enter your message: \n"))
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("Error: ", err)
			break
		}
		handleMessage(conn, message)
	}
}

func handleMessage(conn net.Conn, message string) {
	message = normalizeInput(message)
	log.Printf("Message recieved: %s\n", message)
	if message == "load" {
		//key
		conn.Write([]byte(fmt.Sprintf("Enter key: \n")))
		key, _ := bufio.NewReader(conn).ReadString('\n')
		key = normalizeInput(key)
		value := load(key)
		conn.Write([]byte(fmt.Sprintf("Value: %s \n", value)))
	} else if message == "store"{
		//key
		conn.Write([]byte(fmt.Sprintf("Enter key: \n")))
		key, _ := bufio.NewReader(conn).ReadString('\n')
		key = normalizeInput(key)
		//value
		conn.Write([]byte("Enter value: \n"))
		value, _ := bufio.NewReader(conn).ReadString('\n')
		value = normalizeInput(value)
		store(key, value)
		conn.Write([]byte("Stored\n"))
	} else if message == "search"{
		//pattern
		conn.Write([]byte(fmt.Sprintf("Enter pattern: \n")))
		pattern, _ := bufio.NewReader(conn).ReadString('\n')
		result := search(normalizeInput(pattern))
		conn.Write([]byte(fmt.Sprintf("Found: %s \n", result)))
	} else {
		conn.Write([]byte("Unkown command: "+message+"\n"))
	}
}

func normalizeInput(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}

func load(key string) []byte {
	log.Printf("Recieved: %s", key)
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", DIRNAME, key))
	if err != nil {
		log.Printf("Error: %s", err)
	}
	log.Printf("Loaded: %s", file)
	return file
}

func store(key string, value string) {
	log.Printf("Recieved: %s:%s", key, value)
	file, err := os.Create(fmt.Sprintf("%s/%s", DIRNAME, key))
	defer file.Close()
	if err!=nil {
		log.Print(err)
	}
	file.WriteString(value)
	log.Printf("Stored: %s:%s", key, value)
}

func search(pattern string) string{
	log.Printf("Pattern recieved: %s", pattern)
	file, _ := os.Open(DIRNAME)
	defer file.Close()
	list, _ := file.Readdirnames(0)
	resultArray := []string{}
	for _, name := range list {
		if strings.Contains(name, pattern) {
			resultArray = append(resultArray, name)
		}
	}
	result := strings.Join(resultArray,", ")+"\n"
	log.Printf("Pattern found: %s", result)
	return result
}