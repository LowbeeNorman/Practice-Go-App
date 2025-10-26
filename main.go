package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	const port = ":38759"

	// Start listening on the specified port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("New connection from", conn.RemoteAddr())

		// Handle each client connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Received: %s\n", message)

		// Echo the message back to the client
		_, err := conn.Write([]byte("Echo: " + message + "\n"))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}


	if err := scanner.Err(); err != nil {
		fmt.Println("Connection error:", err)
	}

	fmt.Println("Client disconnected:", conn.RemoteAddr())

}