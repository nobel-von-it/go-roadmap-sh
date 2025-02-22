package main

import (
	"context"
	"log"
	"time"
)

func hello(c chan struct{}) {
	defer close(c)
	for i := 0; i < 10; i++ {
		log.Println("Hello, World!")
		time.Sleep(time.Millisecond * 200)
	}
}

func wait() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	c := make(chan struct{})
	go hello(c)
	select {
	case <-ctx.Done():
		log.Println("timeout")
	case <-c:
		log.Println("done")
	}
}

func main() {
	wait()
}
