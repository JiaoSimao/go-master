package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-master/db"
)

type Student struct {
	ID    int `gorm:"AUTO_INCREMENT"`
	Name  string
	Age   int
	Grade string
}

func main() {
	//假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
	dsn := "root:jsm1234@tcp(192.168.159.132:3306)/homework?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	if err := db.InitDB(dsn); err != nil {
		panic("数据初始化错误：" + err.Error())
	}

	if db.DB == nil {
		panic("数据库未初始化")
	}

	initTableErr := db.DB.AutoMigrate(&Student{})
	if initTableErr != nil {
		panic("failed to init table Student")
	}
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	stu := Student{Name: "张三", Age: 20, Grade: "三年级"}
	result := db.DB.Debug().Create(&stu)
	if result.Error != nil {
		panic("插入学生失败：" + result.Error.Error())
	}

	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	stus := &[]Student{}
	db.DB.Where("age > ?", 18).Find(stus)
	jsonData, jsonErr := json.MarshalIndent(stus, "", "\t")
	if jsonErr != nil {
		panic("json转换错误：" + jsonErr.Error())
	}
	fmt.Println(string(jsonData))

	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	stu3 := Student{}
	db.DB.Debug().Model(&stu3).Where("name = ?", "张三").Update("grade", "四年级")

	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	//创建一条小于15岁的学生
	student := &Student{Name: "李四", Age: 14, Grade: "二年级"}
	db.DB.Debug().Create(student)
	////删除小于15岁的学生记录
	db.DB.Debug().Model(&Student{}).Where("age < ?", 15).Delete(&Student{})

}
