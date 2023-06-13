package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var R *redis.Client

func InitRedis() {
	R = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6380",
		Password: "",
	})

	res, err := R.Ping().Result()
	fmt.Println(res, err)
}

func main() {
	InitRedis()
	fmt.Println(R.Get("hh"))
}
