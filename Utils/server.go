package server

import "net"

// StartServer - starts a server
func StartServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		_, err := ln.Accept()
		if err != nil {
			// handle error
		}
		// go handleConnection(conn)
	}
}
