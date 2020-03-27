package srv

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
		srv.run()
	}()

	dl := net.Dialer{Timeout: time.Second * 20}
	wg := sync.WaitGroup{}
	gcnt := 4098
	for i := 0; i < gcnt; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := dl.Dial("tcp", srvAddr)
			if err != nil {
				fmt.Printf("dial err: %v", err)
				return
			}
			defer conn.Close()
			_, err = conn.Write([]byte("hello"))
			if err != nil {
				fmt.Printf("write to conn err: %v", err)
				return
			}
			time.Sleep(time.Second * 3)
			_, err = conn.Write([]byte("stop"))
			if err != nil {
				fmt.Printf("write to conn err: %v", err)
				return
			}
		}()
	}
	fmt.Print("CLIEN LAUNCHED")

	wg.Wait()
	srv.wg.Wait()
	fmt.Println("COUNT SRV CONN", srv.connCnt)
}
