package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("New HTTP connection from %s '%s'", r.RemoteAddr, r.Header.Get("User-Agent"))
	fmt.Fprintf(w, Taco)
}

func httpServer() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Printf("failed to listen on 80 (using port 8000 instead): %s", err)
		err := http.ListenAndServe(":8000", nil)
		if err != nil {
			log.Fatal("failed to listen on 8000: %s", err)
		}
	}
}
