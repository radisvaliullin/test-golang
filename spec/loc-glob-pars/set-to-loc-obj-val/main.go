package main

import (
	"log"

	"github.com/radisvaliullin/test-golang/spec/loc-glob-pars/set-to-loc-obj-val/pack"
)

func main() {

	o := pack.O()
	log.Printf("o - %+v", o)

	o.A = "aa"
	log.Printf("o - %+v", o)
}
