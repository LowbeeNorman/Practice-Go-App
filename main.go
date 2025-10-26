package main

import (
    "bufio"
    "fmt"
    "net"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()
    clientAddr := conn.RemoteAddr().String()
    fmt.Println("Client connected:", clientAddr)

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        text := scanner.Text()
        fmt.Printf("Received from %s: %s\n", clientAddr, text)
        fmt.Fprintln(conn, "Echo:", text)
    }

    fmt.Println("Client disconnected:", clientAddr)
}

func main() {
    listener, err := net.Listen("tcp", ":38759") // listens on all interfaces
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server listening on port :38759")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}
