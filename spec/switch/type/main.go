package main

import (
	"fmt"
)

func main() {

	var i interface{}
	s := "someText"

	i = s

	switch t := i.(type) {
	case string:
		fmt.Println("is string - ", t)
	default:
		fmt.Println("unknown type ", t)
	}
}
