package main

import (
	"log"
	"reflect"
)

type I interface {
	One()
	Two(int) error
}

func main() {

	t := reflect.TypeOf((*I)(nil)).Elem()
	var s []string
	for i := 0; i < t.NumMethod(); i++ {
		s = append(s, t.Method(i).Name)
	}

	log.Printf("I methods - %+v", s)
}
