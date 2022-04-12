package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3724"
	CONN_TYPE = "tcp"
)

var LOGIN_CHALL = byte(0x00)
var LOGIN_PROOF = byte(0x01)
var REALMLIST = byte(0x10)

func main() {
	log.Print("Starting login serve")
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		log.Printf("Accepting connection from: %v", conn.LocalAddr())
		// Handle connections in a new goroutine.
		go handleRequest(conn)

	}
}

func handleRequest(conn net.Conn) {
	var buf [1024]byte
	header := make([]byte, 4)

	for {
		n, err := conn.Read(header)
		//log.Print("Here is n")
		//log.Print(n)
		if n != 4 || err != nil {
			log.Print("Bad")
		}

		// Calculate the packet size
		size := 0
		size = size + int(header[0])
		size = size + int(header[1])*256

		// Allocate the appropriate size for our data (size - 2 bytes used for the length
		data := make([]byte, size-2)

		n, err = conn.Read(buf[0:])
		if err != nil {
			return
		}

		//fmt.Println(buf[0:])
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}

		log.Print(data)
	}
}
