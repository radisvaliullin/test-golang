package main

import "log"

func main() {

	// map
	var m map[int]int
	log.Printf("m is %v", m)
	// m[1] = 1
	// log.Printf("m is %v", m)

	// slice
	var s []int
	log.Printf("s is %v", s)
	s = append(s, 1)
	log.Printf("s is %v", s)
}
