package main

import (
	"fmt"
	"time"
)

func main() {
	//初始化一个int型的channel，长度为10
	messages := make(chan int, 10)
	done := make(chan bool)

	//程序退出的时候close channel
	defer close(messages)

	//consumer消费者
	//起一个新的协程
	go func() {
		//节拍器，每一秒从管道获取一个事件
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			//1.如果done了，执行
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("receive message: %d\n", <-messages)
			}
		}
	}()
	//producer生产者
	for i := 0; i < 10; i++ {
		messages <- i
	}
	//主线程睡5秒
	time.Sleep(5 * time.Second)
	//关闭done通道，这样在消费者里监听done关闭了，就会退出
	close(done)
	time.Sleep(1 * time.Second)
	//打印主程序退出
	fmt.Println("main process exit!")
}
