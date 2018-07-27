package main

import "log"

func main() {

	var sl []int
	log.Printf("sl - %v", sl)

	for _, s := range sl {
		log.Printf("inter empty sl, s - %v", s)
	}
}
