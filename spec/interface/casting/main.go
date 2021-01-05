package main

import "fmt"

type x interface {
	x() int
}

type xy interface {
	x() int
	y() int
}

type xyObj struct {
	xx int
	yy int
}

func (o *xyObj) x() int {
	return o.xx
}

func (o *xyObj) y() int {
	return o.yy
}

func main() {

	o := xyObj{xx: 1, yy: 2}

	var ix x
	var ixy xy
	var ixy2 xy

	ixy = &o
	ix = ixy

	// can't casting back to more specialised interface
	// ixy2 = ix

	// use assert
	ixy2 = ix.(xy)

	fmt.Println(ixy2)
}
