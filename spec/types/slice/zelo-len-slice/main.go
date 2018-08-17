package main

import "fmt"

func main() {

	var sl0 []byte
	sl1 := []byte{}
	sl2 := []byte{1}
	fmt.Printf("%v : %v : %v\n", sl0, sl1, sl2)

	if sl0 == nil {
		fmt.Println("sl0 is nil")
	}
	if sl1 == nil {
		fmt.Println("sl1 is nil")
	}
	if sl2 == nil {
		fmt.Println("sl2 is nil")
	}
}
