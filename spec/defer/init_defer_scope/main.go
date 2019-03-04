package main

import "log"

type ob struct {
	msg string
}

func (o *ob) getMsg() string {
	return o.msg
}

func main() {

	var sl []int

	sl = []int{1, 2, 3}
	defer func() {
		log.Printf("def 1, sl is %v\n", sl)
	}()

	sl = nil
	defer func() {
		log.Printf("def 2, sl is %v\n", sl)
	}()

	var po *ob

	po = &ob{msg: "po one"}
	defer po.getMsg()
	defer func() {
		if po != nil {
			po.getMsg()
		} else {
			log.Printf("po one is nil\n")
		}
	}()

	po = nil
	defer func() {
		if po != nil {
			po.getMsg()
		} else {
			log.Printf("po two is nil\n")
		}
	}()
}
