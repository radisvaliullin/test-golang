package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	sum := uint64(0)
	smax := sum
	out := false

	for i := 1; i <= 1000000; i++ {

		is := strconv.Itoa(i)

		for _, s := range is {

			n, err := strconv.Atoi(string(s))
			if err != nil {
				log.Fatal("atoi s, ", err)
			}

			sum += uint64(n)

			if sum < smax {
				out = true
			}
			smax = sum
		}
	}

	fmt.Printf("sum - %v; out - %v", sum, out)
}
