package main

import (
	"fmt"
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

func main() {

	cln, err := redis.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		log.Fatal("redis dial err ", err)
	}

	res := cln.Cmd("set", "test3", "zcxv")
	fmt.Println(res)
}
