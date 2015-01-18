package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func telnetServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:25")
	if err != nil {
		log.Printf("failed to listen on 25 (listening on 2500 instead): %s", err)
		listener, err = net.Listen("tcp", "0.0.0.0:2500")
		if err != nil {
			log.Fatal("failed to listen on 2500: %s", err)
		}
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept telnet connection: %s", err)
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()

			log.Printf("New telnet connection from %s", conn.RemoteAddr())
			fmt.Fprintf(conn, strings.Replace(Taco, "\n", "\r\n", -1))
		}(conn)
	}
}
