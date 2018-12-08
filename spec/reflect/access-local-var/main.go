package main

import (
	"log"
	"reflect"

	"github.com/radisvaliullin/test-golang/spec/reflect/access-local-var/pack"
)

func main() {

	o := pack.O()

	v := reflect.ValueOf(*o)
	log.Printf("o value of - %v", v)

	e := v.FieldByName("b")
	log.Printf("o elem 1 - %v", e)
}
