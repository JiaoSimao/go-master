package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (employee *Employee) PrintInfo() {
	fmt.Printf("员工姓名：%s,员工年龄：%d, 员工工号：%s ", employee.Person.Name, employee.Person.Age, employee.EmployeeID)

}

func main() {
	//使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
	//组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

	emp := Employee{
		EmployeeID: "100",
		Person: Person{
			Name: "张三",
			Age:  20,
		},
	}
	emp.PrintInfo()
}
