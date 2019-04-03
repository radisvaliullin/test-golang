package main

import (
	"log"

	"github.com/radisvaliullin/test-golang/stdlib/runtime/simple-trace/trace"
)

func main() {

	trace.Trace()

	log.Print("end")
}
