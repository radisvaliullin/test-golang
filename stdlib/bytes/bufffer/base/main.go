package main

import (
	"bytes"
	"fmt"
)

func main() {

	buf := bytes.NewBuffer(make([]byte, 0, 8))
	fmt.Println(buf.Len(), buf.Cap())

	buf.Write([]byte{32, 1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(buf.Len(), buf.Cap())

	fmt.Println(buf.Bytes())
}
