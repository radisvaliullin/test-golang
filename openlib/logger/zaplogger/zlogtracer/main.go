package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {

	zlog, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("zap logger init err, %v", err)
	}

	// main log
	zlog.Error("test error")

	// goroutine log
	go func() {
		zlog.Debug("go test debug")
		zlog.Error("go test error")
	}()

	//
	for {
	}
}
