package trace

import (
	"log"

	"github.com/go-errors/errors"
)

func Trace() {
	e := errors.New("test errr")
	TracePrint(e)
	// est := e.ErrorStack()
	// log.Printf("error stack - %v", est)
}

func TracePrint(e *errors.Error) {
	// e = errors.New("test errr 2")
	log.Print("trace print")
	est := e.ErrorStack()
	log.Printf("error stack - %v", est)
}
