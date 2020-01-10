package main

import (
	"fmt"
	"math"
	"runtime"
)

// try messure number of thread used by go runtime (using os tools).
func main() {

	// set gomaxprocs
	mp := 8
	fmt.Printf("gomaxprocs old - %v, new - %v\n", runtime.GOMAXPROCS(mp), mp)

	// number of goroutine
	var gcnt = 1000
	// iter count
	var n = 100000
	// delay count num
	var sec = int(math.Pow(10.0, 9.0))

	// start goroutine
	for i := 0; i < gcnt; i++ {

		go func(id int) {

			for i := 0; i < n; i++ {
				fmt.Printf("goroutine %v tick, %v\n", id, i)
				for j := 0; j < sec; j++ {
				}
			}
		}(i)
	}
	fmt.Printf("%v goroutines runned\n", gcnt)

	for i := 0; i < n; i++ {
		fmt.Printf("main tick, %v\n", i)
		for j := 0; j < sec; j++ {
		}
	}
}
