package main

import (
	"log"
	"sync/atomic"
)

func main() {

	var i int32

	i = 33
	log.Printf("i - %v", i)

	atomic.CompareAndSwapInt32(&i, 1, 42)
	log.Printf("i - %v", i)

	atomic.CompareAndSwapInt32(&i, 33, 42)
	log.Printf("i - %v", i)

}
