package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		log.Println("Waiting for a client to connect")

		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go do(conn)
	}
}

func do(conn net.Conn) {
	log.Println("processing the request")
	conn.Write([]byte("+PONG\r\n"))
	conn.Close()
}
