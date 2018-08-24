package main

import "fmt"

func main() {

	m := map[int][]int{}
	fmt.Printf("m - %+v, mp - %p\n", m, m)

	m[0] = append(m[0], 0)
	fmt.Printf("m - %+v, mp - %p, m[0]p - %p, m[0][0]p - %p\n", m, m, m[0], &(m[0][0]))

	m[0] = append(m[0], 0)
	fmt.Printf("m - %+v, mp - %p, m[0]p - %p, m[0][0]p - %p\n", m, m, m[0], &(m[0][0]))

	fmt.Printf("end, m - %+v\n", m)

	// m2
	m2 := map[int][]int{}
	fmt.Printf("m2 - %+v, m2p - %p\n", m2, m2)
	m2[0] = make([]int, 0, 10)
	fmt.Printf("m2 - %+v, m2p - %p, m2[0]p - %p\n", m2, m2, m2[0])

	m2[0] = append(m2[0], 0)
	fmt.Printf("m2 - %+v, m2p - %p, m2[0]p - %p, m2[0][0]p - %p\n", m2, m2, m2[0], &(m2[0][0]))

	m2[0] = append(m2[0], 0)
	fmt.Printf("m2 - %+v, m2p - %p, m2[0]p - %p, m2[0][0]p - %p\n", m2, m2, m2[0], &(m2[0][0]))
	fmt.Printf("m2 - %+v, m2p - %p, m2[0]p - %p, m2[0][1]p - %p\n", m2, m2, m2[0], &(m2[0][1]))

	fmt.Printf("end, m2 - %+v\n", m2)
}
