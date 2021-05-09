package main

import (
	"fmt"
	"time"
)

/**
 * 参数定义为单向通道，只能发，不能收
 * 通道由发送方关闭
 */
func sendData(ch chan<- int) {
	ch <- 1
	fmt.Printf("[%s] send element to channel ch1: %v\n",
		time.Now().Format("2006-01-02 15:04:05.000"), 1)
	time.Sleep(time.Second)
	ch <- 2
	fmt.Printf("[%s] send element to channel ch1: %v\n",
		time.Now().Format("2006-01-02 15:04:05.000"), 2)
	time.Sleep(time.Second)
	ch <- 3
	fmt.Printf("[%s] send element to channel ch1: %v\n",
		time.Now().Format("2006-01-02 15:04:05.000"), 3)
	// 三秒后关闭通道
	time.AfterFunc(time.Second*3, func() {
		close(ch)
		fmt.Printf("[%s] close channel ch1.\n",
			time.Now().Format("2006-01-02 15:04:05.000"))
	})
}


func main() {
	// 声明并初始化了一个元素类型为int、容量为3的通道ch1
	ch1 := make(chan int, 3)

	// 异步发送数据，然后关闭通道
	go sendData(ch1)

	// 从通道接收元素值的时候，同样要用接送操作符<-，只不过，这时需要把它写在变量名的左边，用于表达“要从该通道接收一个元素值”的语义。
	for {
		elem, ok := <-ch1
		if !ok {
			//通道关闭则结束循环
			fmt.Printf("[%s] The channel ch1 is closed.\n", time.Now().Format("2006-01-02 15:04:05.000"))
			break
		}
		fmt.Printf("[%s] received element from channel ch1: %v\n",
			time.Now().Format("2006-01-02 15:04:05.000"), elem)
	}

}
