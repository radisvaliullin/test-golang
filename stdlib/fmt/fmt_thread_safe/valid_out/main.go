package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/radisvaliullin/test-golang/stdlib/fmt/fmt_thread_safe/bigslices"
)

func main() {
	fmt.Println("main start")

	f := os.Stdin

	cntA, cntB := 0, 0

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, bigslices.Size)
	scanner.Buffer(buf, bigslices.Size*2)
	for scanner.Scan() {

		line := scanner.Bytes()
		if len(line) != (bigslices.Size - 1) {
			log.Printf("wrong line length - %v\n", len(line))
		}
		if bytes.Compare(bigslices.BigSlice0[:bigslices.Size-1], line) == 0 {
			cntA += 1
		}
		if bytes.Compare(bigslices.BigSlice1[:bigslices.Size-1], line) == 0 {
			cntB += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("line scan err: %v\n", err)
	}
	fmt.Printf("cntA - %v : cntB - %v  :  writeCycles - %v\n", cntA, cntB, bigslices.WriteCycles)
	if cntA != bigslices.WriteCycles || cntB != bigslices.WriteCycles {
		log.Printf("a and b lines count wrong")
	}
	fmt.Println("main end")
}
