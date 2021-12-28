package main

import (
	"fmt"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
		fmt.Printf("produce data: %d\n", i)
	}
	close(ch)
}

func costomer(ch <-chan int) {
	for k := range ch {
		fmt.Printf("Get data:%d\n", k)
	}
}

func main() {
	ch := make(chan int, 5)
	go produce(ch)
	costomer(ch)
}
