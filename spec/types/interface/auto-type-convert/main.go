package main

import "log"

func main() {

	some(nil)
}

func some(a interface{}) {

	switch a.(type) {
	case int:
		log.Printf("is int")
	case []int:
		log.Printf("is slice of int")
	case []interface{}:
		log.Printf("is interface{} slice")
	default:
		log.Printf("UNKNOWN")
	}
}
