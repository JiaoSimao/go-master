package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter2 struct {
	num int64
}

func (counter *Counter2) Add() {
	atomic.AddInt64(&counter.num, 1)
}

func (counter *Counter2) Get() int64 {
	return atomic.LoadInt64(&counter.num)
}

func main() {
	//使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
	var wg sync.WaitGroup
	counter := &Counter2{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Add()
			}
		}()

	}
	fmt.Println(counter.Get())
}
