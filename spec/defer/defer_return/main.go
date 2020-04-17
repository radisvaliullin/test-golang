package main

import "fmt"

func main() {

	i := deferTest()
	fmt.Printf("i - %v\n", i)
	i2 := deferTest2()
	fmt.Printf("i - %v\n", i2)
	pi := deferTest3()
	fmt.Printf("i - %v\n", *pi)
}

func deferTest() int {
	var i int
	var pi *int
	pi = &i
	defer func() {
		*pi = 42
	}()
	i = 37
	return i + 1
}

func deferTest2() int {
	var i int
	defer func() {
		i = 42
	}()
	i = 37
	return i + 1
}

func deferTest3() *int {
	var i int
	var pi *int
	pi = &i
	defer func() {
		*pi = 42
	}()
	*pi = 37
	return pi
}
