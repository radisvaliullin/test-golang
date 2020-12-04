package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	var memprofile = flag.String("memprofile", "", "write cpu profile to file")

	flag.Parse()
	if *cpuprofile != "" {
		log.Println("start cpu profiling")
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		runtime.SetCPUProfileRate(10000)
		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal("start profiling:", err)
		}
		defer pprof.StopCPUProfile()
	}
	if *memprofile != "" {
		log.Println("start memory profiling")
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		err = pprof.WriteHeapProfile(f)
		if err != nil {
			log.Fatal("start memory profiling:", err)
		}
		defer pprof.StopCPUProfile()
	}

	o := struct {
		one int
	}{}
	o.one = 33

	time.Sleep(time.Second * 2)
	fmt.Println(o)
}
