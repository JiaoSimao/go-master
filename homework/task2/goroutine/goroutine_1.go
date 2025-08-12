package main

import (
	"fmt"
	"time"
)

func jiShu() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("打印从1到10的奇数：", i)
		}
	}
}

func ouShu() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("打印从2到10的偶数：", i)
		}
	}
}

func main() {
	//题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	go jiShu()
	go ouShu()

	time.Sleep(1 * time.Second)
}
