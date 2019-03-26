package main

import (
	"errors"
	"log"
	"strings"
)

type er struct {
	msg       string
	parentErr error
}

func ner(e error) error {
	return er{parentErr: e}
}

func (e er) Error() string {
	return e.stackPrint(e)
}

func (e er) stackPrint(err error) string {
	eStr := ""
	parErrStr := ""
	if ero, ok := err.(er); ok {
		eStr = ero.msg
		parErrStr = e.stackPrint(ero.parentErr)
		eStr = strings.Join([]string{eStr, parErrStr}, " : ")
	} else {
		eStr = err.Error()
	}
	return eStr
}

func main() {

	e := errors.New("one")

	e2o := er{parentErr: e, msg: "two"}
	var e2 error = e2o

	e3o := er{parentErr: e2, msg: "free"}
	var e3 error = e3o
	log.Printf("e3 - %v, e3 p - %p, e3o p - %p", e3, &e3, &e3o)

	e3o2 := er{parentErr: e3, msg: "wrapped e3"}
	e3 = e3o2
	log.Printf("e3 - %v, e3 p - %p, e3o2 p - %p", e3, &e3, &e3o2)
	if e3o22, ok := e3.(er); ok {
		log.Printf("e3 - %v, e3 p - %p, e3o22 p - %p", e3, &e3, &e3o22)
	} else {
		log.Printf("e3 - %v, e3 p - %p, e3 is NOT er", e3, &e3)
	}

	e3 = e3
	log.Printf("e3 - %v, e3 p - %p", e3, &e3)
	if e3val, ok := e3.(er); ok {
		log.Printf("e3 - %v, e3 p - %p, e3val p - %p, e3val t - %T", e3, &e3, &e3val, e3val)
	} else {
		log.Printf("e3 - %v, e3 p - %p, e3 value is NOT er", e3, &e3)
	}
	if e3val, ok := e3.(error); ok {
		log.Printf("e3 - %v, e3 p - %p, e3val p - %p, e3val t - %T", e3, &e3, &e3val, e3val)
	} else {
		log.Printf("e3 - %v, e3 p - %p, e3 value is NOT er", e3, &e3)
	}
}
