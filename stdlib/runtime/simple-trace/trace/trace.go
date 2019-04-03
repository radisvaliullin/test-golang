package trace

import (
	"log"
	"runtime"
	"runtime/debug"
)

func Trace() {

	buf := make([]byte, 1024)
	i := runtime.Stack(buf, false)
	log.Print("stack i ", i)

	log.Printf("buf - %v", string(buf[:i]))
}

func Trace2() {
	log.Print("trace2")
	buf := debug.Stack()

	log.Printf("buf - %v", string(buf))
}
