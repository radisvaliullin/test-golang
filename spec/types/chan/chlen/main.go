package main

import "log"

func main() {

	ch := make(chan int, 5)
	log.Printf("ch len - %v", len(ch))

	ch <- 1
	ch <- 2
	log.Printf("ch len - %v", len(ch))
}
