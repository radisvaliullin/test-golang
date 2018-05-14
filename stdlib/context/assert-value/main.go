package main

import (
	"context"
	"log"
)

type CtxKey string

func main() {

	ctx := context.WithValue(context.Background(), CtxKey("testKey"), "testValue")

	log.Println("ctx testKey - ", ctx.Value(CtxKey("testKey")).(string))

	ek, ok := ctx.Value(CtxKey("emptyKey")).(string)
	if ok {
		log.Println("ctx empty key - ", ek)
	} else {
		log.Println("ctx empty key not found")
	}

	// if skip will panic for empty key
	log.Println("ctx testKey - ", ctx.Value(CtxKey("emptyKey")).(string))
}
