package main

import (
	"fmt"
	"regexp"
)

func main() {

	samples := []string{
		"asdfaf 12345678",
		"as@#$%^&*()*&)(*&)^dfaf 12345678",
		" asdfaf 12345678",
		"asdfaf 12345678 ",
		"    ",
	}

	pattern := `^\S.+\S$`
	var rc = regexp.MustCompile(pattern)

	for _, s := range samples {
		isMatch := rc.MatchString(s)

		if isMatch {
			fmt.Printf("%s is Match pattern %s\n", s, pattern)
		} else {
			fmt.Printf("\"%s\" is NOT match pattern %s\n", s, pattern)
		}
	}
}
