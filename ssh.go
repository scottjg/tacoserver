package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

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

				fmt.Fprintf(channel, strings.Replace(Taco, "\n", "\r\n", -1))
				channel.Close()
			}
		}(conn)
	}
}
