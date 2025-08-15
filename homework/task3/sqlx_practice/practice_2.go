package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	//假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	//要求 ：
	//定义一个 Book 结构体，包含与 books 表对应的字段。
	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

	//创建连接
	dsn := "root:jsm1234@tcp(192.168.159.132:3306)/sqlxhomework?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	sqlxDb, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("数据初始化错误：" + err.Error())
	}
	defer sqlxDb.Close()

	//创建表
	createDataBaseSql := `
		CREATE TABLE IF NOT EXISTS book(
			   id INT AUTO_INCREMENT PRIMARY KEY ,
			   title VARCHAR(255) NOT NULL,
			   author VARCHAR(255) NOT NULL,
			   price DECIMAL(10,2) NOT NULL
		)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
       `

	exec, err := sqlxDb.Exec(createDataBaseSql)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(exec.RowsAffected())

	//创建表记录
	books := []Book{
		{Title: "水浒传", Author: "施耐庵", Price: 60},
		{Title: "西游记", Author: "吴承恩", Price: 70},
		{Title: "三国演义", Author: "罗贯中", Price: 40},
		{Title: "红楼梦", Author: "曹雪芹", Price: 30},
	}
	insertSql := `INSERT INTO book(title, author, price) VALUES (:title, :author, :price)`
	_, insertErr := sqlxDb.NamedExec(insertSql, books)
	if insertErr != nil {
		fmt.Println(insertErr.Error())
		return
	}
	//查询
	bookPrice := 50.0
	var bs []Book
	selectSql := `SELECT * FROM book WHERE price > ? order by price desc limit 10`
	selectErr := sqlxDb.Select(&bs, selectSql, bookPrice)
	if selectErr != nil {
		if errors.Is(selectErr, sql.ErrNoRows) {
			fmt.Println("没有找到超过50元的书籍")
			return
		}
		fmt.Println(selectErr.Error())
		return
	}
	for _, b := range bs {
		fmt.Printf("书籍名称：%s, 作者：%s, 价格：%.2f\n", b.Title, b.Author, b.Price)
	}

}
