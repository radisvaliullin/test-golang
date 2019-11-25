package main

import (
	"encoding/json"
	"log"
)

type exampleBase struct {
	One    int    `json:"one"`
	OneStr string `json:"one_str"`
}

//
type example struct {
	exampleBase
	IntNull  *int `json:"int_null"`
	IntNull2 *int `json:"int_null_2"`
	IntNull3 *int `json:"int_null_3,omitempty"`
}

func main() {

	i := 66

	e := example{
		IntNull:  &i,
		IntNull2: nil,
		IntNull3: nil,
	}

	bt, err := json.Marshal(&e)
	if err != nil {
		log.Fatal("marshal err: ", err)
	}
	log.Printf("example: %v", string(bt))
}
