package main

import (
	"fmt"
	"time"
)

func executeTask(ch chan struct{}, task func()) {
	for i := 0; i < 3; i++   {
		<-ch
		task()
		time.Sleep(1 * time.Second)
		ch <- struct{}{}
	}
}

/**
启动4个goroutine，利用channel按顺序循环打印1，2，3，4
*/
func main() {
	ch := make(chan struct{}, 1)

	/**
	 利用延时，让4个goroutine依次进入ch 的读队列
	  每个goroutine执行依次任务后，向通道写入一个值，让读队列的下一个goroutine执行，自己则再次进入读队列排队
	 */
	go executeTask(ch, func() {
		fmt.Print("1")
	})
	time.Sleep(1 * time.Second)
	go executeTask(ch, func() {
		fmt.Print("2")
	})
	time.Sleep(1 * time.Second)
	go executeTask(ch, func() {
		fmt.Print("3")
	})
	time.Sleep(1 * time.Second)
	go executeTask(ch, func() {
		fmt.Print("4")
	})

	ch <- struct{}{}


	//  让程序挂起
	ch1 := make(chan struct{})
	<-ch1
}
