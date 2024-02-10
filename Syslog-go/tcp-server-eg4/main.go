package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST = "localhost"
	PORT = "5044"
	TYPE = "tcp"
)

func main() {

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// incoming request
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	// write data to response
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("Your message poimal is: %v. Received time: %v", string(buffer[:]), time)
	conn.Write([]byte(responseStr))

	// close conn
	conn.Close()
}
