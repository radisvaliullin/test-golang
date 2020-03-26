package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(1)

	go prt()
	delay()
	go prt()
	delay()
	go prt()
	delay()
	go prt()

	for {
		delay()
		go prt()
	}
}

func prt() {
	fmt.Println("goroutine")
}

func delay() {
	i := 0
	for {
		if i > 1000_000_000 {
			break
		}
		i++
	}
}
