package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {

	sl := []byte("qwertyasdfzxcv\n1234567890")

	rd := bytes.NewReader(sl)

	rdr := bufio.NewReaderSize(rd, 256)

	_, err := rdr.Discard(4)
	if err != nil {
		fmt.Println("discard err: ", err)
	}
	out, err := rdr.ReadSlice('\n')
	if err != nil {
		fmt.Println("read slice err: ", err)
	}

	fmt.Println("out: ", string(out))
}
