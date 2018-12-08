package main

import "log"

type iface interface{}

func main() {

	var i1 iface
	var i2 iface
	log.Printf("i1 - %v, pI1 - %v", i1, &i1)
	log.Printf("i2 - %v, pI2 - %v", i2, &i2)

	i := 73
	pi := &i
	log.Printf("i - %v, pI - %v, ppI - %p", i, &i, &i)
	log.Printf("pi - %v, pPI - %v, ppPI - %p", pi, &pi, pi)

	i1 = pi
	i2 = pi
	log.Printf("i1 - %v, pI1 - %v, ppI1 - %p", i1, &i1, i1)
	log.Printf("i2 - %v, pI2 - %v", i2, &i2)

	*pi++
	log.Printf("i - %v, pI - %v, ppI - %p", i, &i, &i)
	log.Printf("pi - %v, pPI - %v, ppPI - %p", pi, &pi, pi)

	t, ok := i2.(*int)
	if ok {
		log.Printf("ok true, t - %v", *t)
	} else {
		log.Printf("ok false")
	}
}
