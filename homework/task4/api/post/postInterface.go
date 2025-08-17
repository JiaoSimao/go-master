package post

import (
	"github.com/gin-gonic/gin"
	"github.com/go-master/task4/db"
	"github.com/go-master/task4/model"
	"net/http"
	"strconv"
)

// AddPost 创建文章
func AddPost(c *gin.Context) {
	//实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
	var postRequest model.PostRequest
	//检查用户传入参数是否符合
	if err := c.ShouldBind(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var post model.Post
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"error": "用户未认证或认证过期"})
		return
	}
	//加强逻辑，userID是any类型，判断下是否是uint类型
	if uid, ok := userID.(uint); ok {
		post = model.Post{
			UserID:  uid,
			Title:   postRequest.Title,
			Content: postRequest.Content,
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID类型转换错误"})
		return
	}

	if err := db.DB.Debug().Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "文章创建成功",
		"post_id": post.ID,
		"title":   post.Title,
	})
}

// GetAllPost 获取所有文章列表
func GetAllPost(c *gin.Context) {
	//获取所有文章列表
	var total int64 //总条数
	// 查询总条数
	db.DB.Table("posts").Joins("left join users on posts.user_id = users.id").Where("posts.deleted_at IS NULL").Count(&total)

	var postResponses []model.PostResponse
	if err := db.DB.Debug().Table("posts").Select("users.username, posts.id, posts.title, posts.user_id").
		Joins("left join users on posts.user_id = users.id").
		Where("posts.deleted_at IS NULL").
		Scopes(db.Paginate(c)).Scan(&postResponses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":      postResponses,
		"total":     total,
		"page":      c.DefaultQuery("page", "1"),
		"page_size": c.DefaultQuery("page_size", "10"),
	})
}

// GetPostDetail 获取文章详情
func GetPostDetail(c *gin.Context) {
	postID, exists := c.GetQuery("postID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请提供文章的ID"})
		return
	}
	//判断是否传了文章ID，按照文章ID查看详情
	postId, err := strconv.Atoi(postID)
	if err != nil {
		// 转换失败（如字符串不是合法数字）
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章ID转换错误"})
		return
	}
	if postId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请提供正确的文章ID"})
		return
	}

	var postResponse model.PostResponse
	if err := db.DB.Debug().Table("posts").Select("users.username, posts.id, posts.title, posts.user_id, posts.content").
		Joins("left join users on posts.user_id = users.id").
		Where("posts.id = ? and posts.deleted_at IS NULL", postId).Scan(&postResponse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": postResponse,
	})
}

// UpdatePost 更新文章
func UpdatePost(c *gin.Context) {
	var postRequest model.PostRequest
	if err := c.ShouldBind(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if postRequest.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请提供文章的ID"})
		return
	}
	//只有文章的作者才能更新
	var post model.Post
	if err := db.DB.Debug().First(&post, postRequest.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"error": "用户未认证或认证过期"})
		return
	}
	//加强逻辑，userID是any类型，判断下是否是uint类型
	if uid, ok := userID.(uint); ok {
		if post.UserID != uid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "非文章本作者不可以修改文章"})
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID类型转换错误"})
		return
	}

	post.Title = postRequest.Title
	post.Content = postRequest.Content
	if err := db.DB.Debug().Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章修改失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "文章修改成功",
		"post_id": post.ID,
		"title":   post.Title,
	})
}

// DeletePost 更新文章
func DeletePost(c *gin.Context) {
	var postRequest model.PostRequest
	if err := c.ShouldBind(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if postRequest.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请提供文章的ID"})
		return
	}
	//只有文章的作者才能删除
	var post model.Post
	//检查是否存在该文章
	if err := db.DB.Debug().First(&post, postRequest.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"error": "用户未认证或认证过期"})
		return
	}
	//加强逻辑，userID是any类型，判断下是否是uint类型
	if uid, ok := userID.(uint); ok {
		if post.UserID != uid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "非文章本作者不可以删除文章"})
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID类型转换错误"})
		return
	}

	if err := db.DB.Debug().Delete(&post, post.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "文章删除成功",
	})
}
