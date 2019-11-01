package main

import (
	"encoding/json"
	"log"
)

func main() {

	i := 66

	example := struct {
		IntNull  *int `json:"int_null"`
		IntNull2 *int `json:"int_null_2"`
		IntNull3 *int `json:"int_null_3,omitempty"`
	}{
		IntNull:  &i,
		IntNull2: nil,
		IntNull3: nil,
	}

	bt, err := json.Marshal(&example)
	if err != nil {
		log.Fatal("marshal err: ", err)
	}
	log.Printf("example: %v", string(bt))
}
