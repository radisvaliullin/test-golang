package main

import (
	"context"
	"log"
	"time"
)

func main() {

	// ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	dtm := time.Now().Add(time.Second * 5)
	ctx, _ := context.WithDeadline(context.Background(), dtm)

	for {
		if err := ctx.Err(); err != nil {
			log.Fatalf("ctx err is %v", err)
		}
		log.Print("sleep")
		time.Sleep(time.Second)
	}
}
