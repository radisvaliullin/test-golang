package main

import "fmt"

func main() {

	if i := 0; i > 0 {
		fmt.Printf("i > 0, i is %v\n", i)
	} else {
		fmt.Printf("i not > 0, i is %v\n", i)
	}
}
