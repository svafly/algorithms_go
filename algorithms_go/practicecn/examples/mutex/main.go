package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//起三个线程
	go lock()
	go rlock()
	go wlock()
	//主程序睡眠5秒
	time.Sleep(5 * time.Second)
}

func lock() {
	//锁
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		//程序退出的时候关闭锁
		defer lock.Unlock()
		fmt.Println("lock:", i)
	}
}
func rlock() {
	//读写分离锁
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		//读锁不会互斥，可以执行
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("rlock:", i)
	}
}
func wlock() {
	//写锁互斥，会被等待
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("wlock:", i)
	}
}
