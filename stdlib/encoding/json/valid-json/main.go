package main

import (
	"encoding/json"
	"log"
)

func main() {

	jl := []string{
		"{}",
		`{"a":1,"b":"2"}`,
		`{asd}`,
		`{1}`,
		`{1,}`,
	}

	for _, j := range jl {

		st := struct{}{}

		if err := json.Unmarshal([]byte(j), &st); err != nil {
			log.Printf("json %v is NOT valid\n", j)
		} else {
			log.Printf("json %v is Valid\n", j)
		}
	}
}
