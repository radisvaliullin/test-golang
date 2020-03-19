package main

import (
	"fmt"
	"time"
)

func main() {

	tzero := time.Time{}
	tzero2 := time.Time{}

	tnow := time.Now()
	tnow2 := tnow
	tnow3 := time.Now()

	if tzero == tzero2 {
		fmt.Println("tzero is equal Time{}")
	}
	if tzero == tnow {
		fmt.Println("tzero is equal time.Now")
	}

	if tnow == tnow2 {
		fmt.Println("tnow is eq tnow2")
	}
	if tnow == tnow3 {
		fmt.Println("tnow is eq tnow3")
	}
	fmt.Println("tnow sub tnow3 ", tnow.Sub(tnow2))
	if tnow.Sub(tnow2) == 0 {
		fmt.Println("tnow sub tnow3 is 0")
	}
}
