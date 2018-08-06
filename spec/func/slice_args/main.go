package main

import "fmt"

func main() {

	sl := []string{"qwerty", "asdf", "zxcv"}
	pnt(sl...)
}

func pnt(args ...string) {
	fmt.Printf("%v\n", args)
}
