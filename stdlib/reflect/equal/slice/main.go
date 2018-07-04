package main

import (
	"log"
	"reflect"
)

func main() {

	sl1 := []string{"one", "two"}
	sl2 := []string{"one", "two"}
	sl3 := []string{"one", "two", " "}

	if reflect.DeepEqual(sl1, sl2) {
		log.Print("sl1 and sl2 is deep equal")
	} else {
		log.Print("sl1 and sl2 is NOT deep equal")
	}

	if reflect.DeepEqual(sl1, sl3) {
		log.Print("sl1 and sl3 is deep equal")
	} else {
		log.Print("sl1 and sl3 is NOT deep equal")
	}
}
