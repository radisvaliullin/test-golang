package main

import (
	"fmt"
	"regexp"
)

func main() {

	sample := "asdfaf 12345678"
	var re = regexp.MustCompile(`([0-9]{3})([0-9]{3})([0-9]{4})`)
	// var re = regexp.MustCompile(`(^|[^_])\bproducts\b([^_]|$)`)
	s := re.ReplaceAllString(sample, "$1 - $2 - $3")

	fmt.Println("res - ", s)
}
