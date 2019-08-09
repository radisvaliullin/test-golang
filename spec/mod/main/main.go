package main

import (
	"log"

	"github.com/radisvaliullin/test-golang/spec/mod/onetwo"
	onetwoV2 "github.com/radisvaliullin/test-golang/v2/spec/mod/onetwo"
)

func main() {

	log.Printf("onetwo v1; %v", onetwo.One())
	log.Printf("onetwo v2: %v", onetwoV2.One())
	log.Printf("onetwo v2: %v", onetwoV2.OneTwo())
}
