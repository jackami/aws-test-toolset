package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestRedisConnPool(t *testing.T) {
	pool := &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "127.0.0.1:7777")
		},
	}

	for i := 0; i < 500; i++ {
		conn := pool.Get()

		_, err := conn.Do("set", "name", "jamey"+strconv.Itoa(i))
		if err != nil {
			fmt.Println("error redis set")
		}
		name, err := redis.String(conn.Do("get", "name"))
		if err != nil {
			fmt.Println("error redis get")
		}
		fmt.Println("name = " + name)

		//conn.Close()
		//pool.Close()
	}

	time.Sleep(time.Second * 30)
}

func TestRedisConnAlone(t *testing.T) {
	for i := 0; i < 500; i++ {
		conn, err := redis.Dial("tcp", "127.0.0.1:7777")
		if err != nil {
			log.Fatalln(err)
		}
		_, err = conn.Do("SET", "Name","Jeff" + strconv.Itoa(i))
		if err != nil {
			log.Fatalln(err)
		}
		name, err := redis.String(conn.Do("GET", "Name"))
		if err != nil {
			fmt.Println("error redis get")
		}
		fmt.Println("Name = " + name)

		//conn.Close()
	}

	time.Sleep(time.Second * 30)
}

func TestRedisConcurrBench(t *testing.T) {
	var wg sync.WaitGroup = sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		go func(idx int, pWg sync.WaitGroup) {
			wg.Add(1)

			conn, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				log.Fatalln(err)
			}
			_, err = conn.Do("SET", "Name" + strconv.Itoa(idx),"Jeff" + strconv.Itoa(idx))
			if err != nil {
				log.Fatalln(err)
			}

			for {
				name, err := redis.String(conn.Do("GET", "Name" + strconv.Itoa(idx)))
				if err != nil {
					fmt.Println("error redis get")
				}
				fmt.Println("Name" + strconv.Itoa(idx) + " = " + name)

				time.Sleep(3 * time.Second)
			}

			wg.Done()

			//conn.Close()
		}(i, wg)
	}

	wg.Wait()
}