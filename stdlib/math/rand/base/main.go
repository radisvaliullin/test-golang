package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//
	fmt.Println(rand.Int63())
	fmt.Println(rand.Int63())
	fmt.Println(rand.Int63())

	//
	src := rand.NewSource(1)
	rnd := rand.New(src)
	fmt.Println(rnd.Int())
	fmt.Println(rnd.Int())
	fmt.Println(rnd.Int())

	//
	src2 := rand.NewSource(time.Now().UnixNano())
	rnd2 := rand.New(src2)
	fmt.Println(rnd2.Int())
	fmt.Println(rnd2.Int())
	fmt.Println(rnd2.Int())
}
