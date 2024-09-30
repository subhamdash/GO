package main

import (
    "fmt"
    "net"
)

func main() {
    // Listen for incoming connections on port 8080
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer listener.Close()
s
    fmt.Println("Server listening on port 8080")

    // Accept incoming connections
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        defer conn.Close()

        // Handle each connection in a separate goroutine
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())

    // Read from the connection
    buffer := make([]byte, 1024)
    _, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }

    // Echo back the received data
    _, err = conn.Write(buffer)
    if err != nil {
        fmt.Println("Error writing:", err)
        return
    }

    fmt.Println("Data echoed back to client")
}

/*

In the Go net package, net.Listen is a function used to create a listener for a specific network address. 
It returns a net.Listener object which can be used to accept incoming connections. Here's how it works:

net.Listen: This function takes two parameters:

The network type ("tcp", "udp", etc.).
The address to listen on (e.g., "localhost:8080").
It creates a listener for the specified network address and returns a net.Listener object or an error if it fails to listen on the address.

net.Listener: This interface represents a network listener. It has a single method, Accept(), 
which blocks until an incoming connection is available. 
When a connection is received, it returns a net.Conn object representing the connection.

Accept(): This method of net.Listener blocks until a connection is made to the listener. 
When a connection is accepted, it returns a net.Conn object representing the connection.

The Accept() method typically runs in a loop, continuously accepting incoming connections. 
You can handle each connection in a separate goroutine to serve multiple clients concurrently.

net.Conn: This interface represents a network connection. It provides methods for reading and writing data to and from the connection. 
Once a connection is accepted, you can use methods like Read(), Write(), Close(), etc., to communicate with the connected client.

*/
