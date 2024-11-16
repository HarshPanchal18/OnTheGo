package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// listen on port
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for{
		// Accept a connection
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleServerConnection(conn)
	}
}

func handleServerConnection(conn net.Conn) {
	var msg string // receive the message
	err := gob.NewDecoder(conn).Decode(&msg)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received message:", msg)
	}

	conn.Close()
}

func client() {
	// Connect to the server
	conn, err := net.Dial("tcp","127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the message
	msg := "Hello, world!"
	fmt.Println("Sending...:", msg)

	err = gob.NewEncoder(conn).Encode(msg)
	if err != nil { fmt.Println(err) }

	conn.Close()
}