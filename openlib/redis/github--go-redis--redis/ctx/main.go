package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	//

	//
	fmt.Println("dl - ", time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	time.Sleep(time.Second * 4)
	dl, ok := ctx.Deadline()
	fmt.Println("dl - ", dl, "; ok - ", ok)
	defer cancel()

	cln := client.WithContext(ctx)
	bts := make([]byte, 1000)
	for i, _ := range bts {
		bts[i] = 65
	}

	stat := cln.Set("test2", string(bts), 0)
	err = stat.Err()
	if err != nil {
		log.Fatal("redis set err ", err)
	}
	fmt.Println("set test ", stat.Name())

	pong, err = cln.Ping().Result()
	fmt.Println(pong, err)

	time.Sleep(time.Second * 3)
	stat = cln.Set("test", string(bts), 0)
	err = stat.Err()
	if err != nil {
		log.Fatal("redis set err ", err)
	}
	fmt.Println("set test ", stat.Name())
}
