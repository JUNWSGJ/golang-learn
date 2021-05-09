package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
 学习 sync/atomic 包
 sync/atomic标准库包中提供的原子操作
 原子操作是比其它同步技术更基础的操作。
 原子操作是无锁的，常常直接通过CPU指令直接实现。
 事实上，其它同步技术的实现常常依赖于原子操作。
 */

func add1000(){
	var n int32 = 0
	var wg sync.WaitGroup
	for i:=0; i<1000; i++ {
		wg.Add(1)
		go func() {
			n++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("add1000 done, n: %d\n", n)
}

func atomicAdd1000(){
	var n int32 = 0
	var wg sync.WaitGroup
	for i:=0; i<1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&n, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("atomicAdd1000 done, n: %d\n", n)
}


func main() {
	add1000()
	atomicAdd1000()
}