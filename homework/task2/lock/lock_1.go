package main

import (
	"fmt"
	"sync"
)

func main() {
	//编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	var mu sync.Mutex
	var count int
	var wg sync.WaitGroup
	wg.Add(10)

	go add_1(&count, &mu, &wg)
	go add_2(&count, &mu, &wg)
	go add_3(&count, &mu, &wg)
	go add_4(&count, &mu, &wg)
	go add_5(&count, &mu, &wg)
	go add_6(&count, &mu, &wg)
	go add_7(&count, &mu, &wg)
	go add_8(&count, &mu, &wg)
	go add_9(&count, &mu, &wg)
	go add_10(&count, &mu, &wg)
	wg.Wait()
	fmt.Println("最后计数器的值：", count)
}

func add_1(num *int, s *sync.Mutex, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		s.Lock()
		*num++
		s.Unlock()
	}
}

func add_2(num *int, s *sync.Mutex, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		s.Lock()
		*num++
		s.Unlock()
	}
}

func add_3(num *int, s *sync.Mutex, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		s.Lock()
		*num++
		s.Unlock()
	}
}

func add_4(num *int, s *sync.Mutex, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		s.Lock()
		*num++
		s.Unlock()
	}
}

func add_5(num *int, s *sync.Mutex, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		s.Lock()
		*num++
		s.Unlock()
	}
}

func add_6(num *int, s *sync.Mutex, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		s.Lock()
		*num++
		s.Unlock()
	}
}

func add_7(num *int, s *sync.Mutex, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		s.Lock()
		*num++
		s.Unlock()
	}
}

func add_8(num *int, s *sync.Mutex, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		s.Lock()
		*num++
		s.Unlock()
	}
}

func add_9(num *int, s *sync.Mutex, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		s.Lock()
		*num++
		s.Unlock()
	}
}

func add_10(num *int, s *sync.Mutex, s2 *sync.WaitGroup) {
	defer s2.Done()
	for i := 1; i <= 1000; i++ {
		s.Lock()
		*num++
		s.Unlock()
	}
}
