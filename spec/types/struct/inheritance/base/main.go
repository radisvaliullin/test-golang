package main

import "fmt"

type a struct {
	a string
}

func (a *a) msg() string {
	return a.a
}

type aa struct {
	a
	aa string
}

type aaa struct {
	a
	aaa string
}

func (a *aaa) msg() string {
	return a.aaa
}

func main() {

	va := a{a: "a"}
	vaa := aa{a: va, aa: "aa"}
	vaaa := aaa{a: va, aaa: "aaa"}

	fmt.Println(vaa.msg())
	fmt.Println(vaaa.msg())
}
