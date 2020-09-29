package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	t := time.Unix(1591888922635544035/int64(time.Second), 1591888922635544035%int64(time.Second))
	tt := time.Unix(0, 1591888922635544035)
	fmt.Println(t)
	fmt.Println(tt)

	t2, err := time.Parse("2006-01-02 15:04:05 -0700 UTC", "2020-06-11 15:26:11 +0000 UTC")
	if err != nil {
		log.Fatalf("time parse error %v\n", err)
	}
	fmt.Println(t2)
	fmt.Println(t2.Unix())
	fmt.Println("1591889177822920785")
}
