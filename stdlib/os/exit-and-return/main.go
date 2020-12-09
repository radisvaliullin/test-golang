package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		fmt.Println("defer")
	}()

	fmt.Println("befoer os exit")
	os.Exit(1)
	fmt.Println("after os exit")

	return
}
