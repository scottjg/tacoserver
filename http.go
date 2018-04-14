package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	log.Printf("New HTTP connection from %s '%s'", r.RemoteAddr, userAgent)

	if strings.Contains(userAgent, "Mozilla") {
		fmt.Fprint(w, metaRefreshHTML(NextTaco()))
		return
	}

	fmt.Fprintf(w, NextTaco())
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
