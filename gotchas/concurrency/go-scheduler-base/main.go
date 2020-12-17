package main

import "fmt"

func main() {

	n := 10
	ch := make(chan int, n+1)

	for i := 0; i < n; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}
	ch <- 73

	for i := range ch {
		fmt.Println(i)
	}
}
