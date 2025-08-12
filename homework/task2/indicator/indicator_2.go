package main

import "fmt"

func main() {
	//实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	var a []int = []int{1, 2, 3, 4, 5}

	mul(&a)
	fmt.Println(a)

}

func mul(i *[]int) {
	for index := range *i {
		(*i)[index] *= 2
	}
}
