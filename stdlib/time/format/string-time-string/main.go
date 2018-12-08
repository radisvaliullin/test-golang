package main

import (
	"log"
	"time"
)

func main() {

	timeStr := "2018-05-15T19:00:00+03:00"

	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		log.Fatalf("time parse err - %v", err)
	}
	log.Printf("time is %v", t)
	log.Printf("time is %v", t.UTC())

	log.Printf("time is %v", t.Format(time.RFC3339))
	log.Printf("time is %v", t.UTC().Format(time.RFC3339))
}
