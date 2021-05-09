package main

import (
	"fmt"
	"time"
)

var channels = [3]chan int{
	nil,
	make(chan int),
	nil,
}



func main(){
	intChan := make(chan int, 1)
	// 一秒后关闭通道。
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}

