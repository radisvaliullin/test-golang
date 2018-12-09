package main

import (
	"log"
	"sync/atomic"
	"unsafe"
)

func main() {

	ob := obj{}
	ob.rt.Store((*rt)(nil))
	log.Print("ob - ", ob.rt)

	v := value{v: 64}
	log.Printf("v - %v", v)
	log.Printf("v, vp - %p", &v)
	log.Printf("v.v, v.vp - %p", &v.v)
	vp := (*ifaceWords)(unsafe.Pointer(&v))
	log.Printf("vp, vpp - %p", vp)
	log.Printf("%T", vp)
	log.Printf("vp, typ - %v, typT - %T", vp.typ, vp.typ)
	log.Printf("vp, data - %v", vp.data)

	var x interface{}
	log.Printf("x - %v, xp - %p", x, &x)
	x = 73
	xp := (*ifaceWords)(unsafe.Pointer(&x))
	log.Printf("xp, xpp - %p", xp)
	log.Printf("%T", xp)
	log.Printf("xp, typ - %v, typT - %T", xp.typ, xp.typ)
	log.Printf("xp, data - %v", xp.data)

	var ip *int
	ip = (*int)(xp.data)
	log.Printf("ip, ipv - %v", *ip)
}

type obj struct {
	rt atomic.Value
}

type rt struct {
	max float64
}

// ifaceWords is interface{} internal representation.
type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

type value struct {
	v interface{}
}
