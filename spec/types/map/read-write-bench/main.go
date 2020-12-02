package main

import (
	"fmt"
	"time"
)

func main() {

	map32 := map[string]interface{}{}

	map1024 := map[string]interface{}{}

	map100_000 := map[string]interface{}{}

	map1000_000 := map[string]interface{}{}

	kn := 1000_000
	keys := genKeys(kn)

	// fill maps
	for i := 0; i < kn; i++ {
		k := keys[i]
		if i < 32 {
			map32[k] = k
		}
		if i < 1024 {
			map1024[k] = k
		}
		if i < 100_000 {
			map100_000[k] = k
		}
		if i < 1000_000 {
			map1000_000[k] = k
		}
	}

	key32 := keys[:32]
	key1024 := keys[500:532]
	key100_000 := keys[50_000:50_032]
	key1000_000 := keys[500_000:500_032]

	begin := time.Now()
	for _, k := range key32 {
		map32[k] = k
	}
	end := time.Now()
	fmt.Println("write to map32:", end.Sub(begin))

	begin = time.Now()
	for _, k := range key1024 {
		map1024[k] = k
	}
	end = time.Now()
	fmt.Println("write to map1024:", end.Sub(begin))

	begin = time.Now()
	for _, k := range key100_000 {
		map100_000[k] = k
	}
	end = time.Now()
	fmt.Println("write to map100_000:", end.Sub(begin))

	begin = time.Now()
	for _, k := range key1000_000 {
		map1000_000[k] = k
	}
	end = time.Now()
	fmt.Println("write to map1000_000:", end.Sub(begin))

	time.Sleep(time.Second)
	fmt.Println(len(map32), len(map1024), len(map100_000), len(map1000_000))

}

func genKeys(n int) []string {
	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		k := fmt.Sprintf("%015d", i)
		res = append(res, k)
	}
	return res
}
