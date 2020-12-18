package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan []byte, 10)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(ch chan []byte) {
		defer wg.Done()
		arr := [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
		sl := arr[:]
		// sl := []byte{1, 2, 3, 4, 5, 6, 7, 8}
		fmt.Printf("main: sl last item address %p\n", &(sl[7]))
		ch <- sl
	}(ch)
	wg.Add(1)
	go func(ch chan []byte) {
		defer wg.Done()
		sl := <-ch
		fmt.Printf("thread: sl last item address %p\n", &(sl[7]))
		ch <- sl
		sl[7] = 88
	}(ch)
	wg.Wait()

	sl := <-ch
	fmt.Println(sl)
}
