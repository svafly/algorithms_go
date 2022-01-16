package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var flag = false

//多生产者、消费者模式
func main() {
	ch := make(chan string, 10)
	//消费、生产计数器
	wgPd := new(sync.WaitGroup)
	wgCs := new(sync.WaitGroup)

	//3个生产者
	for i := 0; i < 3; i++ {
		wgPd.Add(1)
		go producer(i, wgPd, ch)
	}
	//2个消费者
	for j := 0; j < 2; j++ {
		wgCs.Add(1)
		go consumer(wgCs, ch)
	}

	//退出生产
	go func() {
		time.Sleep(time.Second * 3)
		flag = true
	}()
	wgPd.Wait()
	close(ch)
	wgCs.Wait()
}

func producer(threadID int, wg *sync.WaitGroup, ch chan string) {
	count := 0
	for !flag {
		time.Sleep(time.Second * 1)
		count++

		data := strconv.Itoa(threadID) + "------" + strconv.Itoa(count)
		fmt.Printf("producer,%s\n", data)
		ch <- data
	}
	wg.Done()
}
func consumer(wg *sync.WaitGroup, ch chan string) {
	for data := range ch {
		time.Sleep(time.Second * 1)
		fmt.Printf("consumer,%s\n", data)
	}
	wg.Done()
}
