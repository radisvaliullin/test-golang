package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"
)

// Example use context package for make non blocking multiple requests
func main() {

	// run test servers
	// servers accept our non blocking requests
	go startTCPServ()

	// init context variable
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	// setup context with timeout, after timeout expiration all request will abort
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	_ = cancel

	// non blocking tcp req (to our test tcp server)
	res, err := tcpReq(ctx)
	if err != nil {
		log.Fatal("call tcpReq: err - ", err)
	}
	log.Printf("call tcpReq res: %v", string(res))

}

// simple response model struct
type respData struct {
	body []byte
	err  error
}

// tcp send request
func tcpReq(ctx context.Context) ([]byte, error) {

	// dial connection with timeout, simple solution
	dial := net.Dialer{Timeout: time.Second * 20}
	conn, err := dial.DialContext(ctx, "tcp", ":7374")
	if err != nil {
		log.Print("tcpReq: dial err - ", err)
		return nil, err
	}

	// response chan
	c := make(chan respData, 1)

	// do async send/resp
	go func() { c <- tcpSendResp(conn) }()

	// schedule timeout
	// if context timeout expired, then kill do async request goroutine
	select {
	case <-ctx.Done():
		// if timeout expired kill request by conn.Close() (thread safe)
		conn.Close()
		<-c // wait restReqResp
		return nil, ctx.Err()
	case rd := <-c:
		// if request success, return response data
		if rd.err != nil {
			return nil, rd.err
		}
		return rd.body, nil
	}
}

// tcp request response async handler
func tcpSendResp(conn net.Conn) respData {
	defer conn.Close()

	// send request data
	send := []byte("longlife")
	_, err := conn.Write(send)
	if err != nil {
		log.Print("tcpSendResp: send err - ", err)
		return respData{err: err}
	}

	// handle req response data
	resp := make([]byte, len("longlife"))
	_, err = io.ReadFull(conn, resp)
	if err != nil {
		log.Print("tcpSendResp: read resp err - ", err)
		return respData{err: err}
	}

	return respData{body: resp}
}

// servers

// simple tcp server
func startTCPServ() {

	// time.Sleep(time.Second * 60)

	// listen
	ln, err := net.Listen("tcp", ":7374")
	if err != nil {
		log.Fatal("tcp srv listen: err - ", err)
	}

	// accept handler
	go func() {

		time.Sleep(time.Second * 60)

		for {

			// accept
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal("tcp srv accept: err - ", err)
			}

			// connection handler
			go func(conn net.Conn) {
				defer conn.Close()

				// read cln request
				req := make([]byte, len("longlife"))
				_, err := io.ReadFull(conn, req)
				if err != nil {
					log.Fatal("tcp srv conn handler: req err - ", err)
				}

				// long life pause
				time.Sleep(time.Second * 60)

				// send req response to client
				_, err = conn.Write(req)
				if err != nil {
					log.Fatal("tcp srv conn handler: resp err - ", err)
				}

			}(conn)
		}
	}()
}
