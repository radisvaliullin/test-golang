package main

import (
	"log"
	"net"
	"sync"
)

func main() {

	log.Printf("start main")

	addr := "0.0.0.0:84"

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen err: %v", err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Printf("accept err: %v", err)
				break
			}

			wg.Add(1)
			go connHandler(conn, &wg)
		}
	}()

	wg.Wait()

	log.Printf("end main")
}

func connHandler(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	buf := make([]byte, 1024)

	for {

		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("conn read err: %v", err)
			break
		}
		if n == 0 {
			log.Printf("conn read 0 message")
		}

		log.Printf("conn read msg: %v", string(buf[:n]))
	}
}
