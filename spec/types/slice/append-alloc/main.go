package main

import "fmt"

func main() {

	sl := make([]byte, 0, 1000)
	fmt.Printf("sl - %v ; len -  %v ; cap - %v\n", sl, len(sl), cap(sl))
	// fmt.Printf("sl - %v ; len -  %v ; cap - %v\n", sl[990:], len(sl), cap(sl))

	for i := 0; i < 1000; i++ {
		sl = append(sl, byte(i))
		// fmt.Printf("sl - %v ; len -  %v ; cap - %v\n", sl, len(sl), cap(sl))
	}
	// fmt.Printf("sl - %v ; len -  %v ; cap - %v\n", sl, len(sl), cap(sl))
	fmt.Printf("sl - %v ; len -  %v ; cap - %v\n", sl[990:], len(sl), cap(sl))

	sl = sl[998:]
	fmt.Printf("sl - %v ; len -  %v ; cap - %v\n", sl, len(sl), cap(sl))

	sl = append(sl, 10)
	fmt.Printf("sl - %v ; len -  %v ; cap - %v\n", sl, len(sl), cap(sl))

	// t := cap(sl) - len(sl)
	// for i := 0; i < t; i++ {
	// 	sl = append(sl, byte(i))
	// 	fmt.Printf("sl - %v ; len -  %v ; cap - %v\n", sl, len(sl), cap(sl))
	// }

	// sl = sl[7:]
	// fmt.Printf("sl - %v ; len -  %v ; cap - %v\n", sl, len(sl), cap(sl))

	// sl = append(sl, 10)
	// fmt.Printf("sl - %v ; len -  %v ; cap - %v\n", sl, len(sl), cap(sl))

}
