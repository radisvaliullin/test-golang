package main

import "context"

func main() {

	ctx := context.Background()

	ctx2, cancel := context.WithCancel(ctx)

	cancel()
	cancel()

	select {
	case <-ctx2.Done():
		println("ctx2 done")
	default:
		println("ctx2 do NOT done")
	}
}
