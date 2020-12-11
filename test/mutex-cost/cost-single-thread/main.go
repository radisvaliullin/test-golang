package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var loopSize = 1_000_000_000

	var inc int
	var now time.Time
	var since time.Duration

	// inc without mutex
	now = time.Now()
	for i := 0; i < loopSize; i++ {
		inc++
	}
	since = time.Since(now)
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

	// fibo
	fb := fibo{}
	now = time.Now()
	for i := 0; i < loopSize; i++ {
		_ = fb.next()
	}
	since = time.Since(now)
	fmt.Println("next fibo loop time:", since)

	// fibo mutex
	fb = fibo{}
	now = time.Now()
	for i := 0; i < loopSize; i++ {
		mx.Lock()
		_ = fb.next()
		mx.Unlock()
	}
	since = time.Since(now)
	fmt.Println("next fibo loop time:", since)
}

type fibo struct {
	prev int
	curr int
}

func (f *fibo) next() int {
	if f.prev == 0 && f.curr == 0 {
		f.curr = 1
		return f.prev
	}
	f.prev, f.curr = f.curr, f.prev+f.curr
	return f.prev
}
