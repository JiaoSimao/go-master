package main

import (
	"fmt"
	"sync"
)

func main() {
	//实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	var wg sync.WaitGroup
	//缓冲10的通道
	ch := make(chan int, 10)
	wg.Add(2)

	go producer(ch, &wg)
	go consumer(ch, &wg)
	wg.Wait()
}

func consumer(ch chan int, s *sync.WaitGroup) {
	defer s.Done()
	for value := range ch {
		fmt.Println("value:", value)
	}
}

func producer(ch chan int, s *sync.WaitGroup) {
	defer s.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}
