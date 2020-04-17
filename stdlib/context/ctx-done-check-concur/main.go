package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/context"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("go 1 done.")
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("go 2 done.")
		}
	}()

	time.Sleep(time.Second * 5)
	cancel()

	wg.Wait()
	fmt.Println("done")
}
