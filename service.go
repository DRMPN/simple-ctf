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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	defer log.Println("Disconnected")
	log.Println("Connection established")
	for{
		log.Println("Waiting for user message")
		conn.Write([]byte("Enter your message: "))
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
	log.Printf("Recieved: %s\n", message)
	if message == "load" {
		//key
		conn.Write([]byte(fmt.Sprintf("Enter key: ")))
		key, _ := bufio.NewReader(conn).ReadString('\n')
		value := load(key)
		conn.Write([]byte(fmt.Sprintf("Value: %s", value)))
	} else if message == "store"{
		//key
		conn.Write([]byte(fmt.Sprintf("Enter key: ")))
		key, _ := bufio.NewReader(conn).ReadString('\n')
		//value
		conn.Write([]byte("Enter value: "))
		value, _ := bufio.NewReader(conn).ReadString('\n')
		store(key, value)
		conn.Write([]byte("Stored\n"))
	} else if message == "search"{
		conn.Write([]byte(fmt.Sprintf("Response: %s\n",message)))
	} else {
		conn.Write([]byte("Unkown command: "+message+"\n"))
	}
}

func normalizeInput(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}

func load(key string) []byte {
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", DIRNAME, key))
	if err != nil {
		log.Printf("Error: %s", err)
	}
	return file
}

func store(key string, value string) {
	file, err := os.Create(fmt.Sprintf("%s/%s", DIRNAME, key))
	defer file.Close()
	if err!=nil {
		log.Print(err)
	}
	file.WriteString(value)
}

func search(pattern string) {
	//TODO
}