package main

import "fmt"

func main() {

	var i Ider
	var i2 Ider

	h := &Header{ID: 77}

	o := &One{Header: Header{ID: 44}}

	i = h
	i2 = &o.Header

	fmt.Println(i.Id(), i2.Id())

}

type Ider interface {
	Id() int
}

type Header struct {
	ID int
}

func (h *Header) Id() int {
	return h.ID
}

type One struct {
	Header Header
	O      int
}

func (o *One) One() int {
	return o.O
}
