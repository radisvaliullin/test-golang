package main

import "fmt"

type iface interface {
	A() string
}

type num int

func (n *num) A() string {
	return fmt.Sprintf("%v", n)
}

type struc struct {
	n int
}

func (s *struc) A() string {
	return fmt.Sprintf("%v", s.n)
}

func main() {

	var i iface

	n := num(33)
	s := struc{n: 34}
	i = &n

	fmt.Printf("%v\n", i)

	switch v := i.(type) {
	case *num:
		fmt.Println("i is num")
	case *struc:
		fmt.Println("i is struc")
	default:
		fmt.Println("i is unknown ", v)
	}

	i = &s
	switch v := i.(type) {
	case *num:
		fmt.Println("i is num")
	case *struc:
		fmt.Println("i is struc")
	default:
		fmt.Println("i is unknown ", v)
	}
}
