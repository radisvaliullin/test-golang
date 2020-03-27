package main

import (
	"log"
	"net/http"
)

func main() {

	log.Printf("start main")

	addr := "0.0.0.0:80"

	mux := http.NewServeMux()
	mux.HandleFunc("/", reqHandler)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("listen and server err: %v", err)
	}
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("req handler")
}
