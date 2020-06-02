package main

import (
	"bytes"
	"fmt"
)

func main() {

	buf := bytes.NewBuffer(make([]byte, 0, 32))
	fmt.Println(buf.Len(), buf.Cap())

	buf.Write([]byte{32})
	fmt.Println(buf.Len(), buf.Cap())

	fmt.Println(buf.Bytes())
}
