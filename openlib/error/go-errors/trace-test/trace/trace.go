package trace

import (
	"log"

	"github.com/go-errors/errors"
)

func Trace() {
	e := errors.New("test errr")
	est := e.ErrorStack()
	log.Printf("error stack - %v", est)
}
