package main

import (
	"fmt"
	"net"
)

func main() {

	l, err := net.Listen("tcp", ":7373")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}

	if err := l.Close(); err != nil {
		fmt.Println("first close err:", err)
	}
	if err := l.Close(); err != nil {
		fmt.Println("second close err:", err)
	}
}
