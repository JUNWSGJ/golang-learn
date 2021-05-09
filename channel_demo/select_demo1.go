package main

import (
	"fmt"
	"math/rand"
)


func main(){
	//这里用一个包含了三个候选分支的select语句，
	//分别尝试从上述三个通道中接收元素值，哪一个通道中有值，哪一个对应的候选分支就会被执行。
	//后面还有一个默认分支，不过在这里它是不可能被选中的。

	// 准备好3个通道。
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送元素值。
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Println("No candidate case is selected!")
	}
}
