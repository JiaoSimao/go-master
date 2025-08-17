package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-master/task4/model"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

// GenerateToken 生成token
func GenerateToken(username string, id uint) (string, error) {
	// 设置过期时间
	expirationTime := time.Now().Add(2 * time.Hour)

	// 创建声明
	claims := &model.Claims{
		UserID:   id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "jwt-auth-app",
		},
	}
	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString([]byte(config.Jwt.JwtSecret))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// JwtAuthMiddleware jwt中间件验证
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头中获取Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供签名令牌"})
			c.Abort()
			return
		}

		//检查令牌格式是否正确
		strArr := strings.Split(authHeader, " ")
		if len(strArr) != 2 && strArr[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请提供正确格式的签名令牌"})
			c.Abort()
			return
		}

		tokenString := strArr[1]
		claims, err := parseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
	}
}

func parseToken(tokenString string) (*model.Claims, error) {
	//解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名方法: %v", token.Header["alg"])
		}
		return []byte(config.Jwt.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效的令牌")
}
