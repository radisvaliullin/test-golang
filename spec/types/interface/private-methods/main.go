package main

import (
	"fmt"
)

type I interface {
	Run()
}

type IO struct {
	Msg string
}

func (o *IO) Run() {
	o.print()
}

func (o *IO) print() {
	fmt.Println("private print: ", o.Msg)
}

func main() {

	io := IO{Msg: "io msg"}
	var i I = &io

	i.Run()
}
