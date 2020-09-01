package main

import (
	"fmt"
	"time"
)

func main() {

	tmr := time.NewTimer(time.Second)

	if tmr.Stop() {
		fmt.Println("stop")
	}
	if !tmr.Stop() {
		fmt.Println("not stop")
	}
}
