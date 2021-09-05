package main

import (
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

func Test_RedisAuth(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.137.18:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		TLSConfig: &tls.Config{

		},
	})

	pong, err := client.Ping(nil).Result()
	fmt.Println(pong, err)

	err = client.Set(nil,"feekey", "examples", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(nil,"feekey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("feekey", val)

	val2, err := client.Get(nil,"feekey2").Result()
	if err == redis.Nil {
		fmt.Println("feekey does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("feekey", val2)
	}
}
