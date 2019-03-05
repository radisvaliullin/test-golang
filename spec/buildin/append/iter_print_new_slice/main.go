package main

import "fmt"

func main() {

	dist := []int{56, 57, 58}
	// dist := make([]int, 4)
	// dist[0] = 56
	// dist[1] = 57
	// dist[2] = 58
	fmt.Println(dist)
	fmt.Printf("%p : %p\n", dist, &(dist[0]))
	fmt.Println(dist[:0], dist[:1], dist[:2])

	newD := append(dist[:0], dist[1:]...)
	fmt.Println(len(dist), cap(dist), dist)
	fmt.Println(len(newD), cap(newD), newD)
	fmt.Printf("%p : %p\n", dist, &(dist[0]))
	fmt.Println(dist[:0], dist[:1], dist[:2])

	// for i := 0; i < len(dist); i += 1 {
	// 	fmt.Println(i, i+1)
	// 	fmt.Println(dist[:i], dist[i+1:])
	// 	nedD := append(dist[:i], dist[i+1:]...)
	// 	fmt.Println(nedD, i)
	// }
}
