package main

import (
	"log"
	"net/http"
)

var addr = "0.0.0.0:7373"

func main() {

	// http server
	srv := &http.Server{Addr: addr}

	// muxer
	mux := http.NewServeMux()

	mux2 := http.NewServeMux()
	// mux2.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Print("mux2 root")
	// })
	mux2.HandleFunc("hand/", mux2Handler)
	// mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Print("/mux2")
	// 	w.WriteHeader(200)
	// 	_, err := w.Write([]byte("Hello, /mux2"))
	// 	if err != nil {
	// 		log.Printf("/mux2 handler, err - %v", err)
	// 		return
	// 	}
	// })
	mux.Handle("/mux2/", mux2)
	mux.HandleFunc("/mux", func(w http.ResponseWriter, r *http.Request) {
		log.Print("/mux")
		w.WriteHeader(200)
		_, err := w.Write([]byte("Hello, /mux"))
		if err != nil {
			log.Printf("/mux handler, err - %v", err)
			return
		}
	})
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Print("/ is root")
	// 	w.WriteHeader(200)
	// 	_, err := w.Write([]byte("Hello, / is root"))
	// 	if err != nil {
	// 		log.Printf("/ is root handler, err - %v", err)
	// 		return
	// 	}
	// })

	srv.Handler = mux
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server serve: err - %v", err)
	}
}

func mux2Handler(w http.ResponseWriter, r *http.Request) {

	log.Printf("/mux2/hand")
	w.WriteHeader(200)
	_, err := w.Write([]byte("Hello, /mux2/hand"))
	if err != nil {
		log.Printf("/mux2/hand, err - %v", err)
		return
	}
}
