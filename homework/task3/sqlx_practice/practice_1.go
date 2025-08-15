package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func main() {
	//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	//要求 ：
	//1.编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	//2.编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

	//创建连接
	dsn := "root:jsm1234@tcp(192.168.159.132:3306)/sqlxhomework?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	sqlxDb, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("数据初始化错误：" + err.Error())
	}
	//创建表
	createDataBaseSql := `
		CREATE TABLE IF NOT EXISTS employee(
			   id INT AUTO_INCREMENT PRIMARY KEY ,
			   name VARCHAR(255) NOT NULL,
			   department VARCHAR(255) NOT NULL,
			   salary DECIMAL(10,2) NOT NULL
		)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
       `

	exec, err := sqlxDb.Exec(createDataBaseSql)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(exec.RowsAffected())

	//创建记录
	employees := []Employee{
		{Name: "张三", Department: "技术部", Salary: 8000},
		{Name: "李四", Department: "财务部", Salary: 7000},
		{Name: "王二", Department: "施工部", Salary: 6000},
	}
	insertDataBaseSql := `INSERT INTO employee(name,department,salary) VALUES (:name,:department,:salary)`
	_, insertErr := sqlxDb.NamedExec(insertDataBaseSql, employees)
	if insertErr != nil {
		panic(insertErr.Error())
	}

	//需求1
	var emps []Employee
	selectDataBaseSql := `SELECT * FROM employee where department = ?`
	selectErr1 := sqlxDb.Select(&emps, selectDataBaseSql, "技术部")
	if selectErr1 != nil {
		fmt.Println(selectErr1.Error())
	}
	fmt.Println(emps)

	//需求2
	var emp []Employee
	maxSalarySql := `SELECT * FROM employee ORDER BY salary DESC limit 1`
	selectErr2 := sqlxDb.Select(&emp, maxSalarySql)
	if selectErr2 != nil {
		fmt.Println(selectErr2.Error())
	}
	fmt.Println(emp)

}
