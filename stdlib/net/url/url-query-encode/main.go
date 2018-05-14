package main

import (
	"log"
	"net/url"
)

func main() {

	// set query par
	u := url.URL{
		Scheme:   "https",
		Host:     "ya.ru",
		Path:     "v1/dosome",
		RawQuery: url.Values{"par1": []string{"oneval"}}.Encode(),
	}
	log.Println("u - ", u.String())

	// set empty query par
	u = url.URL{
		Scheme:   "https",
		Host:     "ya.ru",
		Path:     "v1/dosome",
		RawQuery: url.Values{}.Encode(),
	}
	log.Println("u - ", u.String())
}
