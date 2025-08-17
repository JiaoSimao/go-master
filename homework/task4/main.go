package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-master/task4/api/comment"
	"github.com/go-master/task4/api/post"
	"github.com/go-master/task4/api/user"
	"github.com/go-master/task4/db"
	"github.com/go-master/task4/utils"
)

func main() {
	var err error
	//加载数据库配置文件
	config, err := utils.LoadConfig("config.yaml")
	if err != nil {
		panic("配置文件加载错误：" + err.Error())
	}

	//初始化数据库
	err = db.InitDB(config.Mysql)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	noAuthApi := r.Group("/api")
	{
		noAuthApi.POST("/register", user.Register)
		noAuthApi.POST("/login", user.Login)
	}

	authApi := r.Group("/api")
	authApi.Use(utils.JwtAuthMiddleware())
	{
		authApi.POST("/addPost", post.AddPost)
		authApi.GET("/getAllPost", post.GetAllPost)
		authApi.GET("/getPostDetail", post.GetPostDetail)
		authApi.PUT("/updatePost", post.UpdatePost)
		authApi.DELETE("/deletePost", post.DeletePost)

		authApi.POST("/addComment", comment.AddComment)
		authApi.GET("/getPostAllComments", comment.GetPostAllComments)
	}

	err = r.Run(":8280")
	if err != nil {
		return
	}
}
