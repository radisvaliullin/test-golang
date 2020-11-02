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

	buf2 := bytes.NewBuffer(make([]byte, 8))
	fmt.Println(buf2.Len(), buf2.Cap(), buf2.Bytes())
	buf2.Reset()
	fmt.Println(buf2.Len(), buf2.Cap(), buf2.Bytes())
	// buf2.Truncate(buf2.Cap())
	// fmt.Println(buf2.Len(), buf2.Cap(), buf2.Bytes())
}
