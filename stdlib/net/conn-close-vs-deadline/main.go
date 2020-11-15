package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

var srvAddr = "localhost:7373"

func main() {

	wg := sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())

	srvNextSig := make(chan struct{}, 1)
	toChan := make(chan time.Time, 1)

	ln, err := net.Listen("tcp", srvAddr)
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		ln.Close()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {

			select {
			default:
			case <-ctx.Done():
				return
			}

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

					select {
					default:
					case <-ctx.Done():
						return
					}

					to := <-toChan
					if err := conn.SetDeadline(to); err != nil {
						fmt.Println("set deadline 1, err:", err)
						return
					}
					n, err := conn.Read(buff)
					if err != nil {
						if err, ok := err.(net.Error); ok && err.Timeout() {
							fmt.Println("handler read 1 timeout err OK")
						} else {
							fmt.Println("handler read 1 err:", err)
							return
						}
					}
					fmt.Println("handler read bytes: ", buff[:n])
					srvNextSig <- struct{}{}
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
	toChan <- time.Now()
	<-srvNextSig
	toChan <- time.Now().Add(time.Second * 15)

	if err := conn.Close(); err != nil {
		fmt.Println("client conn close:", err)
		return
	}
	cancel()

	wg.Wait()
}
