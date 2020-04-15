package srvcln

import (
	"bufio"
	"log"
	"net"
	"strings"
	"sync"
)

// Srv test server
type Srv struct {
	Addr string

	wg sync.WaitGroup

	connCnt int
}

func (s *Srv) Run() {

	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Printf("listen err: %v", err)
		return
	}

	for {

		conn, err := ln.Accept()
		s.connCnt++
		if err != nil {
			log.Printf("accept err: %v", err)
			continue
		}

		s.wg.Add(1)
		go connHandler(conn, &s.wg)

	}
}

func connHandler(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	for {
		line, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			// log.Printf("srv: conn: read line err: %v", err)
			return
		}
		if strings.TrimSpace(line) == "stop" {
			return
		}
	}
}
