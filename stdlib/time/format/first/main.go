package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Printf("time - %v\n", now.Format(time.RFC3339))
}
