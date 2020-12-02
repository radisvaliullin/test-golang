package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var (
	redisAddr = "0.0.0.0:6378"

	testStream = "test-stream"
)

func main() {

	rOpts := &redis.Options{
		Addr: redisAddr,
	}
	cln := redis.NewClient(rOpts)
	defer cln.Close()
	pong, err := cln.Ping().Result()
	if err != nil {
		fmt.Println("ping err:", err)
		return
	}
	fmt.Println("pong:", pong)

	var lastID string

	xmsgs, err := cln.XRevRangeN(testStream, "+", "-", 1).Result()
	if err != nil && err != redis.Nil {
		fmt.Println("read last message err:", err)
		return
	}
	if err == redis.Nil {
		fmt.Println("read last message: return nil")
	} else {
		if len(xmsgs) > 0 {
			lastID = xmsgs[0].ID
		} else {
			fmt.Println("read lst message: len 0")
			lastID = "0"
		}
	}

	for {
		xReadArgs := &redis.XReadArgs{
			Streams: []string{testStream, lastID},
			Block:   0,
		}
		xstrms, err := cln.XRead(xReadArgs).Result()
		if err != nil && err != redis.Nil {
			fmt.Println("read new message err:", err)
		}
		if err == redis.Nil {
			fmt.Println("read new message: return nil")
		} else {
			if len(xstrms) > 0 && len(xstrms[0].Messages) > 0 {
				fmt.Println("read new message:", xstrms[0].Messages[0])
				fmt.Println("read new messages:", xstrms)
				lastID = xstrms[0].Messages[0].ID
			} else {
				fmt.Println("read new message: len 0:", xstrms)
			}
		}
	}

	// fmt.Println("done")
}
