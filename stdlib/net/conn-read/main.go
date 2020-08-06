package main

import (
	"log"
	"net"
	"time"
)

func main() {

	srvAddr := "0.0.0.0:7373"

	// start server
	go func() {
		ln, err := net.Listen("tcp", srvAddr)
		if err != nil {
			log.Fatalf("listen err: %v", err)
		}

		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatalf("accept err: %v", err)
			}

			go func(conn net.Conn) {
				defer conn.Close()

				buf := make([]byte, 1024)
				for {
					n, err := conn.Read(buf)
					if err != nil {
						log.Printf("read message err: %v", err)
						return
					}
					log.Printf("read message: len - %v, msg - %v", n, buf[:n])
				}

			}(conn)
		}
	}()

	// start client with delay
	time.Sleep(time.Second)

	conn, err := net.Dial("tcp", srvAddr)
	if err != nil {
		log.Printf("client dial err: %v", err)
		return
	}

	n, err := conn.Write([]byte("hello"))
	if err != nil {
		log.Printf("client write hello err: %v", err)
	}
	log.Printf("client write %v bytes", n)
	n, err = conn.Write([]byte("hello"))
	if err != nil {
		log.Printf("client write hello err: %v", err)
	}
	log.Printf("client write %v bytes", n)

	// delay and exit
	time.Sleep(time.Second)
}
