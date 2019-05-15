package main

import "log"

type one struct {
	a string
	b string
}

func main() {

	o := one{a: "a", b: "b"}
	o2 := o
	log.Printf("o - %+v ; o2 - %+v", o, o2)
}
