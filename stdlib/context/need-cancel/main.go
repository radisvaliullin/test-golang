package main

import (
	"context"
	"runtime"
	"time"
)

func main() {
	сtx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Hour))

	go func() {
		ctx, _ := context.WithCancel(сtx)

		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			ctx, _ := context.WithCancel(ctx)
			ctx = ctx
		}
	}()

	for {
		time.Sleep(time.Second)
		println("main ", runtime.NumGoroutine())
	}
}
