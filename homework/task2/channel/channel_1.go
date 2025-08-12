package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	var ch = make(chan int)
	wg.Add(2)
	//生成
	go add_element(ch, &wg)
	//接收
	go reveive_element(ch, &wg)
	wg.Wait()
}

func reveive_element(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range ch {
		fmt.Println("接收的整数：", value)
	}
}

func add_element(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
