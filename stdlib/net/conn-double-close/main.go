package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var srvAddr = "localhost:7373"

func main() {

	wg := sync.WaitGroup{}
	defer wg.Wait()

	ln, err := net.Listen("tcp", srvAddr)
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer ln.Close()

	go func() {

		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("accept err:", err)
				continue
			}

			wg.Add(1)
			go func(conn net.Conn) {
				wg.Done()
				defer func() {
					if err := conn.Close(); err != nil {
						fmt.Println("hadnler defer close err:", err)
					}
				}()

				buff := make([]byte, 1024)
				for {
					n, err := conn.Read(buff)
					if err != nil {
						fmt.Println("handler read err:", err)
						if err := conn.Close(); err != nil {
							fmt.Println("handler close err:", err)
						}
						return
					}
					fmt.Println("handler read bytes: ", buff[:n])
				}

			}(conn)
		}
	}()

	time.Sleep(time.Millisecond * 1000)

	conn, err := net.Dial("tcp", srvAddr)
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}

	_, err = conn.Write([]byte("qwerty"))
	if err != nil {
		fmt.Println("client write err:", err)
		return
	}

	time.Sleep(time.Millisecond * 1000)

	if err := conn.Close(); err != nil {
		fmt.Println("client conn close:", err)
		return
	}

	time.Sleep(time.Millisecond * 1000)
}
