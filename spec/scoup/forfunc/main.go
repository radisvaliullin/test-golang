package main

import (
	"log"
	"sync"
)

func main() {

	sl := []msg{
		{id: 1, msg: "1"},
		{id: 2, msg: "2"},
		{id: 3, msg: "3"},
		{id: 4, msg: "4"},
		{id: 5, msg: "5"},
	}
	wg := sync.WaitGroup{}
	for _, m := range sl {
		wg.Add(1)
		go func(m *msg) {
			defer wg.Done()
			log.Printf("func m is %+v", m)
		}(&m)
	}
	wg.Wait()
}

type msg struct {
	id  int
	msg string
}
