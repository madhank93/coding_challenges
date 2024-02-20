package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	// Ensure we teardown the server when the program exits
	defer listener.Close()

	fmt.Println("Server is listening on port 6379")

	for {
		log.Println("Waiting for a client to connect")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		// Handle client connection
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// Ensure we close the connection after we're done
	defer conn.Close()

	// Read data
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading request", err.Error())
			return
		}

		log.Printf("Received %d bytes", n)
		log.Printf("Received the following data: %s", string(buf[:n]))

		// Write response data back
		log.Println("Processing the request")
		conn.Write([]byte("+PONG\r\n"))
	}
}
