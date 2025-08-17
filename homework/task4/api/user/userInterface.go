package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-master/task4/db"
	"github.com/go-master/task4/model"
	"github.com/go-master/task4/utils"
	"net/http"
)

// Register 用户注册
func Register(c *gin.Context) {
	var userRequest model.UserRequest
	//检查用户传入参数是否符合
	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//检查用户和邮箱是否存在了
	var existUser model.User
	if err := db.DB.Debug().Where("username = ?", userRequest.Username).First(&existUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "用户已存在"})
		return
	}
	if err := db.DB.Debug().Where("email = ?", userRequest.Email).First(&existUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "邮箱已被注册"})
		return
	}

	//对用户密码进行加密处理
	encryptPassword, err := utils.EncryptPassword(userRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//创建新用户
	user := model.User{
		Username: userRequest.Username,
		Password: encryptPassword,
		Email:    userRequest.Email,
	}

	if err := db.DB.Debug().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "用户注册成功",
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

// Login 用户登录
func Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	//验证登录入参
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//检查用户是否存在
	var existUser model.User
	if err := db.DB.Debug().Where("username = ?", loginRequest.Username).First(&existUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名不存在"})
		return
	}
	if err := utils.CheckPassword(loginRequest.Password, existUser.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码错误"})
		return
	}

	//生成jwt令牌
	token, err := utils.GenerateToken(existUser.Username, existUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成令牌失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "登录成功",
		"token":    token,
		"user_id":  existUser.ID,
		"username": existUser.Username,
	})
}
