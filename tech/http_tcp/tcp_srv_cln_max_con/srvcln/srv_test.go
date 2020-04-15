package srvcln

import (
	"fmt"
	"net"
	"sync"
	"testing"
	"time"
)

var srvAddr = "127.0.0.1:8485"

func Test_Srv(t *testing.T) {

	srv := Srv{Addr: srvAddr}

	go func() {
		srv.Run()
	}()

	dl := net.Dialer{Timeout: time.Second * 5}
	wg := sync.WaitGroup{}
	connCnt := safeCounter{}
	go func() {
		for {
			fmt.Printf("COUNT of CONN: %v\n", connCnt.Cnt())
			time.Sleep(time.Second)
		}
	}()
	clnNum := 10000
	for i := 0; i < clnNum; i++ {
		wg.Add(1)
		connCnt.Inc()
		go func() {
			defer wg.Done()
			defer connCnt.Dec()
			conn, err := dl.Dial("tcp", srvAddr)
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
	fmt.Println("CLIEN LAUNCHED")

	wg.Wait()
	srv.wg.Wait()
	fmt.Println("COUNT SRV CONN", srv.connCnt)
}
