package main

import (
	"fmt"
	"net"
	"net/rpc"

)

type Server struct{}

func (server *Server) Negate(i int64, reply *int64) error {
	*reply = i
	return nil
}

func rpcServer() {
	rpc.Register(new(Server))
	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		fmt.Println(err)
		return
	}

	for{
		conn, err := listener.Accept()
		if err != nil { continue }

		go rpc.ServeConn(conn)
	}
}

func rpcClient() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:9001")
	if err != nil {
		fmt.Println(err)
		return
	}

	var result int64
	err = conn.Call("Server.Negate", int64(999), &result) // invokes the named function, waits for it to complete, and returns its error status.
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("Server.Negate(999) =", result)
	}
}