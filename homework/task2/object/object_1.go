package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	//定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
	//在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	r := Rectangle{width: 10, height: 5}
	r_area := r.Area()
	r_perimeter := r.Perimeter()
	fmt.Println("长方形的面积:", r_area)
	fmt.Println("长方形的周长:", r_perimeter)

	c := Circle{radius: 1.5}
	c_area := c.Area()
	c_perimeter := c.Perimeter()
	fmt.Println("圆的面积:", c_area)
	fmt.Println("圆的周长:", c_perimeter)
}
