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

	var viab iab

	vab := ab{v: 4}
	// vab2 := ab{v: 44}
	viab = &vab
	log.Printf("viab is %+v", viab)

	if v, ok := viab.(ia); ok {
		log.Printf("i is ia")
		log.Printf("i is ia, assert value type %T", v)
	} else {
		log.Printf("i is NOT ia")
	}

	if v, ok := viab.(iabc); ok {
		log.Printf("i is iabc")
		log.Printf("i is iabc, assert value type %T", v)
	} else {
		log.Printf("i is NOT iabc")
	}

	if v, ok := viab.(*ab); ok {
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

	var i interface{}

	i = "qwerty"

	s, ok := i.(string)
	if ok {
		log.Println("is string ", ok, s)
	} else {
		log.Println("is not string ")
	}
}
