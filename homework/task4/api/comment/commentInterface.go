package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/go-master/task4/db"
	"github.com/go-master/task4/model"
	"net/http"
	"strconv"
)

// AddComment 添加评论
func AddComment(c *gin.Context) {
	var commentRequest model.CommentRequest
	// 检查传入参数
	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"error": "用户未认证或认证过期"})
		return
	}
	var comment model.Comment
	//加强逻辑，userID是any类型，判断下是否是uint类型
	if uid, ok := userID.(uint); ok {
		comment = model.Comment{
			UserID:  uid,
			PostID:  commentRequest.PostID,
			Content: commentRequest.Content,
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID类型转换错误"})
		return
	}

	if err := db.DB.Debug().Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":    "评论创建成功",
		"comment_id": comment.ID,
		"post_id":    comment.PostID,
		"user_id":    userID,
	})
}

// GetPostAllComments 获取某个文章的所有评论列表
func GetPostAllComments(c *gin.Context) {
	//获取所有文章列表
	var total int64 //总条数
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

	// 查询总条数
	db.DB.Table("comments").Joins("left join posts on comments.post_id = posts.id").
		Joins("left join users on comments.user_id = users.id").
		Where("comments.post_id = ? and comments.deleted_at is NULL", postId).Count(&total)

	//获取所有评论列表
	var commentResponses []model.CommentResponse
	if err := db.DB.Table("comments").Select("comments.id, comments.content, comments.user_id, users.username, posts.id as post_id,  posts.title").
		Joins("left join posts on comments.post_id = posts.id").
		Joins("left join users on comments.user_id = users.id").
		Where("comments.post_id = ? and comments.deleted_at is NULL", postId).
		Scopes(db.Paginate(c)).
		Scan(&commentResponses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":      commentResponses,
		"total":     total,
		"page":      c.DefaultQuery("page", "1"),
		"page_size": c.DefaultQuery("page_size", "10"),
	})
}
