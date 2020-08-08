package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	// самое большое число
	n := 10
	// сколько надо выбрать
	k := 3
	if k > n {
		panic("huyak")
	}

	sl := make([]int, 0, n)

	for i := 1; i <= n; i++ {
		sl = append(sl, i)
	}

	out := make([]int, 0, k)
	for i := 0; i < k; i++ {
		rnBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(sl))))
		if err != nil {
			panic("huyak")
		}
		rn := int(rnBig.Int64())
		out = append(out, sl[rn])
		sl = append(sl[:rn], sl[rn+1:]...)
	}

	fmt.Println(out)

}
