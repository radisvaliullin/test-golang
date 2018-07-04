package main

import (
	"log"
	"reflect"
)

func main() {

	m1 := map[string][]float64{"one": []float64{0.1}}
	m2 := map[string][]float64{"one": []float64{0.1}}
	m3 := map[string][]float64{"one": []float64{0.1, 0.0}}

	if reflect.DeepEqual(m1, m2) {
		log.Print("m1 and m2 is deep equal")
	} else {
		log.Print("m1 and m2 is NOT deep equal")
	}

	if reflect.DeepEqual(m1, m3) {
		log.Print("m1 and m3 is deep equal")
	} else {
		log.Print("m1 and m3 is NOT deep equal")
	}
}
