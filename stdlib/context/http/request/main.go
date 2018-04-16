package main

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// Test context for net.http package
func main() {

	// run test servers
	// servers accept our non blocking requests
	// startHTTPServ()
	startTCPServ()

	// init context variable
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	// setup context with timeout, after timeout expiration all request will abort
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	_ = cancel

	// non blocking resp api req (to our test http server)
	res, err := restReq(ctx)
	if err != nil {
		log.Fatal("call restReq: err - ", err)
	}
	log.Printf("call restReq res: %v", string(res))
}

// rest api req
func restReq(ctx context.Context) ([]byte, error) {

	// build request
	req, err := http.NewRequest("GET", "http://localhost:7374/longlived", nil)
	if err != nil {
		log.Println("restRec: new req err - ", err)
		return nil, err
	}
	req = req.WithContext(ctx)

	// build client
	tr := &http.Transport{}
	cln := &http.Client{Transport: tr}

	// do req
	resp, err := cln.Do(req)
	if err != nil {
		log.Print("restReq: do req, err - ", err)
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("restReq: read resp body, err - ", err)
		return nil, err
	}

	return respBytes, nil
}

// servers

// simple http rest api handle server
func startHTTPServ() {

	mux := http.NewServeMux()

	// long-lived handler
	mux.HandleFunc("/longlived", func(w http.ResponseWriter, r *http.Request) {
		// long lived, delay
		time.Sleep(time.Second * 60)
		w.Write([]byte("Long lived request"))
	})

	// run server listener
	go func() {
		if err := http.ListenAndServe(":7373", mux); err != nil {
			log.Fatal("http server down: err - ", err)
		}
	}()

}

// simple tcp server (with http handler)
func startTCPServ() {

	// listen
	ln, err := net.Listen("tcp", ":7374")
	if err != nil {
		log.Fatal("tcp srv listen: err - ", err)
	}

	// accept handler
	go func() {

		// time.Sleep(time.Second * 5)

		for {

			// accept
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal("tcp srv accept: err - ", err)
			}

			// connection handler
			go connHandler(conn)
		}
	}()
}

// tcp server, http handler
func connHandler(conn net.Conn) {
	defer conn.Close()

	// read request
	buff := make([]byte, 1024)

	n, err := conn.Read(buff)
	if err != nil {
		log.Printf("srv: conn - %v; err - %v; buff - %v",
			conn.RemoteAddr(), err, string(buff[:n]))
		return
	}

	// send response
	rawResp := "HTTP/1.1 200 OK\r\n" +
		"Content-Length: 6\r\n" +
		"Content-Type: text/plain; charset=utf-8\r\n" +
		"Date: Wed, 19 Jul 1972 19:00:00 GMT\r\n\r\n" +
		"Hello.\r\n"
	resp := []byte(rawResp)

	// long-lived delay
	time.Sleep(time.Second * 1)

	// write response
	n, err = conn.Write(resp)
	if err != nil {
		log.Printf("srv: conn - %v; write resp err - %v", conn.RemoteAddr(), err)
		return
	} else if n != len(resp) {
		log.Printf("srv: conn - %v; write resp err - %v",
			conn.RemoteAddr(), errors.New("write resp n != len(writeData)"))
	}
}
