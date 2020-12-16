package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	defer wg.Wait()

	func() {

		ch := make(chan int, 1)

		wg.Add(1)
		go func() {
			defer wg.Done()

			time.Sleep(time.Second * 2)
			fmt.Println("try write ch")
			ch <- 42

		}()

		fmt.Println("func end")
	}()

	fmt.Println("end")
}
