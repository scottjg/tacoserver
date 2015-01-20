package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func sshServer() {
	config := &ssh.ServerConfig{
		NoClientAuth: true,
	}

	privateBytes, err := ioutil.ReadFile("id_rsa")
	if err != nil {
		log.Fatal("failed to read private key")
	}
	privateKey, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("failed to parse private key")
	}
	config.AddHostKey(privateKey)

	listener, err := net.Listen("tcp", "0.0.0.0:22")
	if err != nil {
		log.Printf("failed to listen on 22 (listening on 2200 instead): %s", err)
		listener, err = net.Listen("tcp", "0.0.0.0:2200")
		if err != nil {
			log.Fatal("failed to listen on 2200: %s", err)
		}
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept ssh connection: %s", err)
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()

			sshConn, chans, _, err := ssh.NewServerConn(conn, config)
			if err != nil {
				log.Printf("failed to handshake: %s", err)
				return
			}

			log.Printf("New SSH connection from %s '%s'", sshConn.RemoteAddr(), sshConn.ClientVersion())
			for newChannel := range chans {
				channel, _, err := newChannel.Accept()
				if err != nil {
					log.Printf("could not accept channel: %s", err)
					continue
				}

				go func(channel ssh.Channel) {
					b := make([]byte, 256)
					for {
						n, err := channel.Read(b)
						if err != nil || (n == 1 && (b[0] == 3 || b[0] == 4)) {
							channel.Close()
							break
						}
					}
				}(channel)

				for {
					_, err := fmt.Fprintf(channel, strings.Replace(Taco, "\n", "\r\n", -1))
					if err != nil {
						break
					}
					time.Sleep(time.Second)
					_, err = fmt.Fprintf(channel, "\033[8A" + strings.Replace(Taco2, "\n", "\r\n", -1))
					if err != nil {
						break
					}
					time.Sleep(time.Second)
					fmt.Fprintf(channel, "\033[8A")
				}
				channel.Close()
			}
		}(conn)
	}
}
