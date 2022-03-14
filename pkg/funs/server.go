package funs

import (
	"log"
	"net"
)

func Server(addr string) {
	serv, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer serv.Close()

	// loop listen
	for {
		conn, err := serv.Accept()
		if err != nil {
			log.Println("Failed to accept conn.", err)
		}
		go handleSession(conn)
	}
}
