package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {

	// slice and reader
	b := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = byte(i + 1)
	}
	fmt.Println("b:", b)
	br := bytes.NewReader(b)

	// buffered reader
	bio := bufio.NewReaderSize(br, 17)

	// read bytes
	b2, err := bio.ReadSlice(16)
	if err != nil {
		fmt.Println("read bytes err:", err)
		return
	}
	fmt.Println("read bytes: ", b2)
	// read bytes
	b2, err = bio.ReadSlice(32)
	if err != nil {
		fmt.Println("read bytes err:", err)
		return
	}
	fmt.Println("read bytes: ", b2)

	// buffer state
	fmt.Println(bio.Buffered())
	fmt.Println(bio.Size())

	// read buffer tail
	b3 := make([]byte, bio.Buffered()+1)
	_, err = bio.Read(b3)
	if err != nil {
		fmt.Println("read bio tail:", err)
		return
	}
	fmt.Println("bio tail:", b3)

	brNext, err := br.ReadByte()
	fmt.Println("br next byte: ", brNext, err)

}
