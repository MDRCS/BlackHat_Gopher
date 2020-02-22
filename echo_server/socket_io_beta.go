package main

//As is customary for most languages, you’ll start by building an echo server
//to learn how to read and write data to and from a socket

import (
	"io"
	"log"
	"net"
)

// echo is a handler function that simply echoes received data.
func echo(conn net.Conn) {

	defer conn.Close()
	// Create a buffer to store received data.
	b := make([]byte, 512)

	//an infinite loopy ensures that the server will continue to listen for connections even after one has been received
	for {
		// Receive data via conn.Read into a buffer.
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}

		if err != nil {
			log.Println("Unexpected error")
			break
		}

		log.Printf("Received %d bytes: %s\n", size, string(b))
		// Send data via conn.Write.
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", ":20080")

	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	log.Println("Listening on 0.0.0.0:20080")
	//an infinite loopy ensures that the server will continue to listen for connections even after one has been received
	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		//Recall from earlier discussions in this section that Conn is both a Reader and
		//a Writer (it implements the Read([]byte) and Write([]byte) interface methods).
		log.Println("Received connection")
		if err != nil {
		log.Fatalln("Unable to accept connection")
	}
		// Handle the connection. Using goroutine for concurrency.
		go echo(conn)
	}
}


// Commands-line to execute the script :
//1- go run socket.io
//2- telnet localhost 20080  //is a client that will create a connection and use the script to send n receive data