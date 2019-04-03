package trace

import (
	"log"
	"runtime"
)

func Trace() {

	buf := make([]byte, 1024)
	i := runtime.Stack(buf, false)
	log.Print("stack i ", i)

	log.Printf("buf - %v", string(buf[:i]))
}
