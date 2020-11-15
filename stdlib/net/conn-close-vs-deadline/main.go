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

	srvReadDone := make(chan struct{})
	srvNextSig := make(chan struct{})
	toChan := make(chan time.Time)
	toChanDone := make(chan struct{})

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

				wg.Add(1)
				go func() {
					defer wg.Done()

					var to time.Time
					for {
						select {
						case <-ctx.Done():
							return
						case to = <-toChan:
						}
						if err := conn.SetDeadline(to); err != nil {
							fmt.Println("set deadline 1, err:", err)
							return
						}
						toChanDone <- struct{}{}
					}
				}()

				buff := make([]byte, 1024)
				for {

					select {
					default:
					case <-ctx.Done():
						return
					}

					fmt.Println("start read")
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
					srvReadDone <- struct{}{}
					<-srvNextSig
				}

			}(conn)
		}
	}()

	time.Sleep(time.Millisecond * 500)

	conn, err := net.Dial("tcp", srvAddr)
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}

	time.Sleep(time.Millisecond * 500)

	fmt.Println("set deadline")
	toChan <- time.Now()
	<-toChanDone
	<-srvReadDone
	fmt.Println("set deadline")
	toChan <- time.Now().Add(time.Second * 15)
	<-toChanDone
	srvNextSig <- struct{}{}

	fmt.Println("write message")
	_, err = conn.Write([]byte("qwerty"))
	if err != nil {
		fmt.Println("client write err:", err)
		return
	}
	<-srvReadDone

	if err := conn.Close(); err != nil {
		fmt.Println("client conn close:", err)
		return
	}
	cancel()
	srvNextSig <- struct{}{}

	wg.Wait()
}
