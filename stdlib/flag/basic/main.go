package main

import (
	"flag"
	"log"
)

func main() {

	version := flag.Int("version", 7373, "test version flag")
	flag.Parse()

	log.Printf("version flag value - %v", *version)
}
