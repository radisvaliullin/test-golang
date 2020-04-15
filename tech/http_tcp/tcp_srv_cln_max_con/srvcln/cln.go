package srvcln

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Cln struct {
	SrvAddr   string
	ClientNum int
}

func (c *Cln) Run() {

	dl := net.Dialer{Timeout: time.Second * 5}
	wg := sync.WaitGroup{}
	connCnt := safeCounter{}
	go func() {
		for {
			fmt.Printf("COUNT of CONN: %v\n", connCnt.Cnt())
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < c.ClientNum; i++ {
		wg.Add(1)
		connCnt.Inc()
		go func() {
			defer wg.Done()
			defer connCnt.Dec()
			conn, err := dl.Dial("tcp", c.SrvAddr)
			if err != nil {
				fmt.Printf("dial err: %v\n", err)
				return
			}
			defer conn.Close()
			_, err = conn.Write([]byte("hello"))
			if err != nil {
				// fmt.Printf("write to conn err: %v\n", err)
				return
			}
			time.Sleep(time.Second * 15)
			_, err = conn.Write([]byte("stop"))
			if err != nil {
				// fmt.Printf("write to conn err: %v\n", err)
				return
			}
		}()
	}
	fmt.Println("CLIENTS LAUNCHED")

	wg.Wait()
}
