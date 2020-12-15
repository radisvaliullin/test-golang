package main

import (
	"errors"
	"fmt"
)

func main() {
	err := defer_test(errors.New("one"))
	fmt.Println("main with err:", err)

	err = defer_test(nil)
	fmt.Println("main without err:", err)
}

func defer_test(inerr error) (err error) {
	defer func() {
		if err != nil {
			fmt.Println("defer error is not empty: ", err)
		}
	}()
	return inerr
}
