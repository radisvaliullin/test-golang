package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	buf := make([]byte, 1024)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		log.Fatalf("io read err: %v\n", err)
	}
	fmt.Printf("stdin - %v\n", string(buf[:n]))
}
