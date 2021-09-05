package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var (
	closeChan = make(chan int, 1)
)

func main() {
	var wg sync.WaitGroup = sync.WaitGroup{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 50; i++ {

		go func(idx int, pWg sync.WaitGroup) {
			wg.Add(1)

			conn, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				log.Fatalln(err)
			}
			_, err = conn.Do("SET", "Name"+strconv.Itoa(idx), "Jeff"+strconv.Itoa(idx))
			if err != nil {
				log.Fatalln(err)
			}

			for {
				select {
				case <-closeChan:
					break
				default:
					{
						name, err := redis.String(conn.Do("GET", "Name"+strconv.Itoa(idx)))
						if err != nil {
							fmt.Println("error redis get")
						}
						fmt.Println("Name" + strconv.Itoa(idx) + " = " + name)

						time.Sleep(time.Duration(r.Intn(3)) * time.Second)
					}
				}
			}
			wg.Done()
			//conn.Close()
		}(i, wg)
	}

	signalCH := InitSignal()
	HandleSignal(signalCH)

	wg.Wait()
}

// InitSignal register signals handler.
func InitSignal() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	return c
}

// HandleSignal fetch signal from chan then do exit or reload.
func HandleSignal(c chan os.Signal) {
	// Block until a signal is received.
	for {
		s := <-c
		log.Fatalln("recive a signal: " + s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			log.Fatalln("recive a signal, close goroutins, others")
			return
		case syscall.SIGHUP:
			log.Fatalln("recive a signal, close goroutins, handup")
			closeChan <- 1
			return
		default:
			log.Fatalln("recive a signal, close goroutins, default")
			return
		}
	}
}
