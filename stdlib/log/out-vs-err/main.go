package main

import "log"

import "os"

func main() {

	log.Print("test log")

	log2 := log.New(os.Stdout, "", 2)
	log2.Print("test log2")

	log3 := log.New(os.Stderr, "", 2)
	log3.Print("test log3")
}
