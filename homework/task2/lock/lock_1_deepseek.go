package main

import (
	"fmt"
	"sync"
)

// 定义一个包含计数器和互斥锁的结构体
type Counter struct {
	mu    sync.Mutex
	value int
}

// 提供一个安全的递增方法
func (c *Counter) Increment() {
	c.mu.Lock()         // 加锁，确保同一时间只有一个协程能访问
	defer c.mu.Unlock() // 确保函数退出时释放锁
	c.value++
}

// 提供一个安全的获取值的方法
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程递增1000次
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终结果
	fmt.Printf("最终计数器值: %d\n", counter.GetValue())
}
