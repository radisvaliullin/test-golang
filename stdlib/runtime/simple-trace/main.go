package main

import (
	"log"

	"github.com/radisvaliullin/test-golang/stdlib/runtime/simple-trace/trace"
)

func main() {

	// trace.Trace()
	// go func() { trace.Trace2() }()
	go func() { trace.Trace3() }()

	for {
	}
	log.Print("end")
}
