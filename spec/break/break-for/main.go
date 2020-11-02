package main

import "fmt"

func main() {

one:
	for i := 0; i < 100; i++ {
		fmt.Println("i ", i)
		const target = 4
		if i > target {
			fmt.Println("i > ", target)
			break one
		}
	}

	fmt.Println("end")
}
