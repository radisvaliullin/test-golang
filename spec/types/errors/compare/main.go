package main

import (
	"errors"
	"fmt"
)

var (
	errOne   = errors.New("error one")
	errOneV2 = errors.New("error one")
)

func main() {

	if errors.Is(errOne, errOneV2) {
		fmt.Printf("errOne is equal to errOneV2\n")
	} else {
		fmt.Println("one is not equal to error oneV2")
	}
}
