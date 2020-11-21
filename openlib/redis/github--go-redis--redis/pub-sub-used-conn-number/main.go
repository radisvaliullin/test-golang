package main

import (
	"fmt"
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

	sub := client.Subscribe("chan1")
	chans := sub.Channel()
	_ = chans
	// m := <-chans
	// fmt.Println("m:", m)

	time.Sleep(time.Second * 3)
	fmt.Println("done")
}
