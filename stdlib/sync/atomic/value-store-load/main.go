package main

import "sync/atomic"

func main() {

	v := atomic.Value{}
	v.Store(64)
	v.Store(73.0)
}
