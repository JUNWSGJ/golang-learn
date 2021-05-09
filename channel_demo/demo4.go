package main

import (
	"fmt"
	"time"
)

func process(timeout time.Duration) bool {
	ch := make(chan bool)
	go func() {
		fmt.Println("begin process...")
		// 模拟处理耗时的业务
		time.Sleep(timeout + time.Second)
		ch <- true // block
		fmt.Println("exit goroutine...")
	}()
	select {
	case <-time.After(timeout):
		fmt.Println("timeout")
		return false
	case result := <-ch:
		return result
	}
}

func main() {
	ok := process(10 * time.Second)
	fmt.Printf("ok: %v\n", ok)
	fmt.Printf("hello world")
	select {}
}
