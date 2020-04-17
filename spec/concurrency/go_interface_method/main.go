package main

import (
	"fmt"
	"time"
)

type ist interface {
	conn()
}

type st struct {
	n int
}

func newSt(n int) ist {
	return &st{n: n}
}

func (s *st) conn() {
	for {
		fmt.Printf("st - %v\n", s.n)
		time.Sleep(time.Second * 20)
	}
}

func main() {

	for i := 0; i < 10; i++ {
		s := newSt(i)
		go s.conn()
	}

	fmt.Println("start")

	for {

	}
}
