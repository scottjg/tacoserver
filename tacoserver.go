package main

func main() {
	go httpServer()
	go telnetServer()
	sshServer()
}
