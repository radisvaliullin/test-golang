package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// Test Pool limit and simultaneous executions
func main() {

	client := redis.NewClient(&redis.Options{
		Addr:        "0.0.0.0:6378",
		Password:    "", // no password set
		DB:          0,  // use default DB
		PoolSize:    10,
		PoolTimeout: time.Second * 600,
	})
	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("ping err:", err)
		return
	}
	fmt.Println("pong:", pong)

	//
	wg := sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(i int) {
			wg.Done()
			stat := client.Set("test_key", fmt.Sprintf("%v", i), 0)
			if err := stat.Err(); err != nil {
				fmt.Println("set err:", err)
				return
			}
		}(i)
	}

	wg.Wait()

	time.Sleep(time.Second * 3)
	fmt.Println("done")
}
