package main

import "log"

type ia interface {
	a()
}

type iab interface {
	a()
	b()
}

type iabc interface {
	a()
	b()
	c()
}

type ab struct {
	v int
}

func (v *ab) a() {
	log.Printf("a is %v", v.v)
}

func (v *ab) b() {
	log.Printf("b is %v", v.v*v.v)
}

func main() {

	var i iab

	i = &ab{v: 4}
	log.Printf("iab is %+v", i)

	if v, ok := i.(ia); ok {
		log.Printf("i is ia")
		log.Printf("i is ia, assert value type %T", v)
	} else {
		log.Printf("i is NOT ia")
	}

	if v, ok := i.(iabc); ok {
		log.Printf("i is iabc")
		log.Printf("i is iabc, assert value type %T", v)
	} else {
		log.Printf("i is NOT iabc")
	}

	if v, ok := i.(*ab); ok {
		log.Printf("i is *ab")
		log.Printf("i is *ab, assert value type %T", v)
	} else {
		log.Printf("i is NOT *ab")
	}

	// compiler error
	// if v, ok := i.(ab); ok {
	// 	log.Printf("i is ab")
	// 	log.Printf("i is ab, assert value type %T", v)
	// } else {
	// 	log.Printf("i is NOT ab")
	// }
}
