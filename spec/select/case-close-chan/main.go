package main

import (
	"fmt"
	"time"
)

func main() {

	stop := make(chan struct{})
	tk := time.NewTicker(time.Second * 2)

	go func() {
		defer fmt.Println("go done")
		for {
			select {
			case t := <-tk.C:
				fmt.Println("tick:", t)
			case <-stop:
				fmt.Println("done")
				return
			}
		}
	}()

	time.Sleep(time.Second * 8)
	close(stop)
	for {
		time.Sleep(time.Second * 4)
		fmt.Println("heartbit")
	}

}
