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

func Trace3() {
	log.Print("trace3")

	traceFuncs := make([]uintptr, 50)
	ci := runtime.Callers(0, traceFuncs)
	traceFuncs = traceFuncs[:ci]

	// for i := 0; i < ci; i++ {
	// 	cf := runtime.FuncForPC(traceFuncs[i])
	// 	file, line := cf.FileLine(traceFuncs[i])
	// 	log.Printf("trace: %v : %v : %v", cf.Name(), file, line)
	// }

	frames := runtime.CallersFrames(traceFuncs)

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		log.Printf("frame func: %v ; file - %v ; line - %v", frame.Function, frame.File, frame.Line)
	}
}
