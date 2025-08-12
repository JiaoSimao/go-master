package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	var count int64
	var wg sync.WaitGroup
	wg.Add(10)

	go increment_1(&count, &wg)
	go increment_2(&count, &wg)
	go increment_3(&count, &wg)
	go increment_4(&count, &wg)
	go increment_5(&count, &wg)
	go increment_6(&count, &wg)
	go increment_7(&count, &wg)
	go increment_8(&count, &wg)
	go increment_9(&count, &wg)
	go increment_10(&count, &wg)
	wg.Wait()
	fmt.Println("最后计数器的值：", count)
}

func increment_1(num *int64, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(num, 1)
	}
}

func increment_2(num *int64, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(num, 1)
	}
}

func increment_3(num *int64, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(num, 1)
	}
}

func increment_4(num *int64, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(num, 1)
	}
}

func increment_5(num *int64, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(num, 1)
	}
}

func increment_6(num *int64, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(num, 1)
	}
}

func increment_7(num *int64, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(num, 1)
	}
}

func increment_8(num *int64, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(num, 1)
	}
}

func increment_9(num *int64, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(num, 1)
	}
}

func increment_10(num *int64, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(num, 1)
	}
}
