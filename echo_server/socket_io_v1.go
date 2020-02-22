package main

import (
	"bufio"
	"log"
	"net"
)

func echo_v1(conn net.Conn) {

	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {

		s, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalln("Unable to read data")
		}

		log.Printf("Read %d bytes: %s", len(s), s)
		log.Println("Writing data")

		if _, err := writer.WriteString(s); err != nil {
			log.Fatalln("Unable to write data")
		}

		writer.Flush()
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
		go echo_v1(conn)
	}
}


// Commands-line to execute the script :
//1- go run socket.io
//2- telnet localhost 20080  //is a client that will create a connection and use the script to send n receive data