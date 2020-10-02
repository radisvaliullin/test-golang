package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("before os exit")
		os.Exit(1)
		fmt.Println("after os exit")
	}()

	wg.Wait()
	fmt.Println("done")
}
