package main

import "fmt"

func main() {

	ch := make(chan int)

	close(ch)

	fmt.Println("read from closed chan 1")
	<-ch
	fmt.Println("read from closed chan 1")
	<-ch
}
