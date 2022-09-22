package main

import (
	"fmt"
	"time"
)

var podChan map[string](chan int)

func main() {
	var c chan int
	c = make(chan int, 0)
	go func() {
		time.Sleep(time.Second * 10)
		<-c
		// fmt.Println(a)
	}()
	c <- 0
	fmt.Println("finish")
}

func Filter(key string) {
	var c chan int
	if len(podChan) == 0 {
		c = make(chan int, 100)
	} else {
		c = podChan[key]
		if c == nil {
			c = make(chan int, 0)
		}
	}
	c <- 0
}
