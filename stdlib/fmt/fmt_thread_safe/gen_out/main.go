package main

import (
	"io"
	"log"
	"os"
	"sync"

	"github.com/radisvaliullin/test-golang/stdlib/fmt/fmt_thread_safe/bigslices"
)

func main() {
	// fmt.Println("main start")

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go fileWriter(0, os.Stdout, bigslices.BigSlice0, wg)
	wg.Add(1)
	go fileWriter(1, os.Stdout, bigslices.BigSlice1, wg)

	wg.Wait()

	// fmt.Println("main end")
}

func fileWriter(id int, w io.Writer, msg []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < bigslices.WriteCycles; i++ {
		n, err := w.Write(msg)
		if err != nil {
			log.Fatalf("%v: write file err: %v\n", id, err)
		}
		if n != bigslices.Size {
			log.Printf("%v: write wrong size, n is %v, msg len is %v\n", id, n, len(msg))
		}
	}
	// log.Printf("%v: write OK\n", id)
}
