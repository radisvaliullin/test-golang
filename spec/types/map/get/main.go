package main

import "log"

func main() {

	// m := map[string]struct{}{"one": struct{}{}, "two": struct{}{}}
	var m map[string]struct{}
	if _, ok := m["one"]; ok {
		log.Print("is ok ", ok)
	} else {
		log.Print("ok is ", false)
	}
}
