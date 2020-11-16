package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var loopSize = 1_000_000_000

	var inc int

	// inc without mutex
	now := time.Now()
	for i := 0; i < loopSize; i++ {
		inc++
	}
	since := time.Since(now)
	fmt.Println("loop time:", since)

	// inc with mutex
	mx := sync.Mutex{}
	inc = 0
	now = time.Now()
	for i := 0; i < loopSize; i++ {
		mx.Lock()
		inc++
		mx.Unlock()
	}
	since = time.Since(now)
	fmt.Println("loop time with mutex cost:", since)
}
