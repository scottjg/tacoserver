package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func telnetServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:23")
	if err != nil {
		log.Printf("failed to listen on 23 (listening on 2300 instead): %s", err)
		listener, err = net.Listen("tcp", "0.0.0.0:2300")
		if err != nil {
			log.Fatal("failed to listen on 2300: %s", err)
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
			for {
				_, err := fmt.Fprintf(conn, strings.Replace(Taco1, "\n", "\r\n", -1))
				if err != nil {
					break
				}
				time.Sleep(time.Second)
				_, err = fmt.Fprintf(conn, "\033[10A" + strings.Replace(Taco2, "\n", "\r\n", -1))
				if err != nil {
					break
				}
				time.Sleep(time.Second)
				fmt.Fprintf(conn, "\033[10A")
			}
		}(conn)

		go func (conn net.Conn) {
			defer conn.Close()
			tmp := make([]byte, 256)
			for {
				n, err := conn.Read(tmp)
				if err != nil {
					break
				}

				// ctrl-c or ctrl-d
				if (n == 5 &&
				   tmp[0] == 255 && tmp[1] == 244 && tmp[2] == 255 &&
				   tmp[3] == 253 && tmp[4] == 6) ||
                                   (n == 1 && tmp[0] == 4) ||
                                   (n == 2 && tmp[0] == 255 && tmp[1] == 236) {
					break
				}
			}
		}(conn)
	}
}
