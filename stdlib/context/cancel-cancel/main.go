package main

import "context"

func main() {

	ctx := context.Background()

	ctx2, cancel := context.WithCancel(ctx)
	ctx3, cancel3 := context.WithCancel(ctx2)

	cancel()
	cancel()
	cancel3()
	cancel3()

	select {
	case <-ctx2.Done():
		println("ctx2 done")
	default:
		println("ctx2 do NOT done")
	}
	select {
	case <-ctx3.Done():
		println("ctx3 done")
	default:
		println("ctx3 do NOT done")
	}
}
