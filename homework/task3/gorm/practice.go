package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-master/db"
	"gorm.io/gorm"
)

type User struct {
	ID        int
	Name      string
	PostCount int `gorm:"default:0"`
	Posts     []Post
}

type Post struct {
	ID            int
	Title         string
	UserID        int
	CommentStatus string `gorm:"default:'无评论'"`
	Comments      []Comment
}

type Comment struct {
	ID      int
	Content string
	PostID  int
}

// AfterCreate 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) AfterCreate(tx *gorm.DB) error {
	//根据post的id去查询有多少评论
	var commentCount int64
	tx.Model(&Comment{}).Where("post_id = ?", p.ID).Count(&commentCount)
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count+?", commentCount)).Error
}

// AfterDelete 为 Post 模型添加一个钩子函数，在文章删除时自动更新用户的文章数量统计字段
func (p *Post) BeforeDelete(tx *gorm.DB) error {
	//根据post的id去查询有多少评论
	var commentCount int64
	tx.Model(&Comment{}).Where("post_id = ?", p.ID).Count(&commentCount)
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count-?", commentCount)).Error
}

// AfterCreate 为 Comment 模型添加一个钩子函数，在评论创建时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (p *Comment) AfterCreate(tx *gorm.DB) error {
	var commentCount int64
	tx.Model(&Comment{}).Where("post_id = ?", p.PostID).Count(&commentCount)
	if commentCount > 0 {
		result := tx.Model(&Post{}).Where("id = ?", p.PostID).Update("comment_status", "有评论")
		if result.Error != nil {
			return result.Error
		}
	} else {
		result := tx.Model(&Post{}).Where("id = ?", p.PostID).Update("comment_status", "无评论")
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// BeforeDelete 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (p *Comment) AfterDelete(tx *gorm.DB) error {
	var commentCount int64
	tx.Model(&Comment{}).Where("post_id = ?", p.PostID).Count(&commentCount)
	if commentCount == 0 {
		result := tx.Model(&Post{}).Where("id = ?", p.PostID).Update("comment_status", "无评论")
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func main() {
	//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Posts （文章）、 Comment （评论）。
	//要求 ：
	//使用Gorm定义 User 、 Posts 和 Comment 模型，其中 User 与 Posts 是一对多关系（一个用户可以发布多篇文章）， Posts 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	//编写Go代码，使用Gorm创建这些模型对应的数据库表。

	dsn := "root:jsm1234@tcp(192.168.159.132:3306)/gormhomework?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	if err := db.InitDB(dsn); err != nil {
		panic("初始化数据库失败" + err.Error())
	}

	if db.DB == nil {
		panic("数据库初始化失败")
	}

	//user 与 post
	err := db.DB.AutoMigrate(&User{}, &Post{})
	if err != nil {
		panic("初始化表失败" + err.Error())
	}

	//Posts 与 Comment
	err1 := db.DB.AutoMigrate(&Comment{})
	if err1 != nil {
		panic("初始化表失败" + err1.Error())
	}

	//基于上述博客系统的模型定义。
	//要求 ：
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	//先创建表记录
	users := []User{
		{
			Name: "张三", Posts: []Post{
				{
					Title: "2025java面试宝典",
					Comments: []Comment{
						{Content: "非常不错"},
					},
				},
			},
		},
		{
			Name: "李四", Posts: []Post{
				{
					Title: "2025python面试宝典",
					Comments: []Comment{
						{Content: "比较专业"},
						{Content: "文章写的比较全面，涵盖了大多数面试问的问题"},
					},
				},
			},
		},
	}
	db.DB.Debug().Create(&users)

	//使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	user := User{}
	db.DB.Debug().Preload("Posts.Comments").Where("name = ?", "张三").Find(&user)
	//db.DB.Debug().Preload("Posts").Preload("Posts.Comments").Where("name = ?", "张三").Find(&user)
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		panic("json转换错误" + err.Error())
	}
	fmt.Println(string(jsonData))

	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	var comments []Comment
	//先查commnet,按照postid分组，得到每个post的评论数
	query := db.DB.Debug().Select("post_id, count(*) as comment_count").Group("post_id").Find(&comments)
	//然后查询post与上述query进行join，得到最多评论的post，按照comment_count进行排序
	post := Post{}
	db.DB.Debug().Select("id, title").Joins("left join (?) as sub on posts.id = sub.post_id", query).Order("sub.comment_count desc").First(&post)
	postData, err := json.MarshalIndent(post, "", "  ")
	if err != nil {
		panic("json转换错误" + err.Error())
	}
	fmt.Println(string(postData))

	//删除评论，看是否会更新文章状态
	// 查询文章当前状态
	var updatedPost Post
	db.DB.Debug().First(&updatedPost, 1)
	fmt.Printf("删除评论前，文章评论状态: %s\n", updatedPost.CommentStatus)

	// 删除第一条评论
	var comment1 Comment
	db.DB.Debug().First(&comment1, "id = ?", 1)
	db.DB.Debug().Delete(&comment1)
}
