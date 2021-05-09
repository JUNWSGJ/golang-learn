package main

import (
	"fmt"
)

func useUninitializedChan() {
	var ch1 chan int
	fmt.Printf("ch: %v", ch1)
	ch1 <- 1 // fatal error: all goroutines are asleep - deadlock!
}

/**
关闭未初始化的chan
panic: close of nil channel
 */
func closeNilChannel() {
	var ch1 chan int
	close(ch1)
}

/**
向已经关闭的chan发送数据
panic: send on closed channel
*/
func sendToClosedChannel() {
	ch1 := make(chan int)
	close(ch1)
	ch1 <- 1
}

/**
关闭已经关闭的chan
panic: close of closed channel
*/
func closeClosedChannel() {
	ch1 := make(chan int)
	close(ch1)
	close(ch1)
}

/**
展示使用chan触发panic的情景
 */
func main() {

	closeNilChannel()
	//sendToClosedChannel()
	//closeClosedChannel()

}
