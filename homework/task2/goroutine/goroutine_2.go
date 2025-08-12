package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	tasks := map[string]func(){
		"task1": func() {
			time.Sleep(100 * time.Millisecond)
		},
		"task2": func() {
			time.Sleep(200 * time.Millisecond)
		},
		"task3": func() {
			time.Sleep(300 * time.Millisecond)
		},
	}

	var wg sync.WaitGroup
	results := make(chan string, len(tasks))

	//启动所有的任务
	for name, task := range tasks {
		wg.Add(1)
		go runTask(name, task, &wg, results)
	}

	go func() {
		//等待任务全部执行完成，关闭通道
		wg.Wait()
		close(results)
	}()

	//打印结果
	for res := range results {
		fmt.Println(res)
	}
}

func runTask(name string, task func(), s *sync.WaitGroup, results chan string) {
	defer s.Done()
	start := time.Now()
	task()
	elapsed := time.Since(start)
	results <- fmt.Sprintf("%s took %s", name, elapsed)
}
