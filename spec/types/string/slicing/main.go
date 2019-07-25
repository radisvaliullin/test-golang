package main

import "log"

func main() {

	st := "1234"

	st2 := st[len(st):]

	log.Printf("st - %v; st2 - %v", st, st2)
}
