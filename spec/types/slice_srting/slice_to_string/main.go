package main

import "fmt"

func main() {

	sl := []byte("abcd")

	str := string(sl)

	fmt.Printf("sl - %v : str - %v\n", sl, str)

	sl[1] = 'z'

	fmt.Printf("sl - %v : str - %v\n", sl, str)
}
