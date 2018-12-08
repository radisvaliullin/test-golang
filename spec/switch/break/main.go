package main

import "log"

func main() {

	log.Print("before switch")
	switch {
	case true:
		break
		log.Print("after break")
	}
	log.Print("after switch")
}
