package main

import "log"

func main() {

	sl := []int{1, 2, 3, 4}

	sl2 := sl[len(sl):]

	log.Printf("sl - %v ; sl2 - %v", sl, sl2)
}
