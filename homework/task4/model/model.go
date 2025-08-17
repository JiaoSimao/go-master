package model

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

type UserRequest struct {
	Username string `json:"username" binding:"required,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	Email    string `form:"email" binding:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PostRequest struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type PostResponse struct {
	Username string `json:"username"`
	ID       uint   `json:"post_id"`
	UserID   uint   `json:"user_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
type CommentRequest struct {
	ID      uint   `json:"id"`
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
}
type CommentResponse struct {
	ID       uint   `json:"id"`
	Content  string `json:"comment_content"`
	UserID   uint   `json:"comment_user_id"`
	Username string `json:"comment_username"`
	PostID   uint   `json:"post_id"`
	Title    string `json:"title"`
}

type Config struct {
	Mysql MysqlConfig `yaml:"mysql"`
	Jwt   JwtConfig   `yaml:"jwt"`
}
type MysqlConfig struct {
	DSN          string `yaml:"dsn"`
	MaxIdleConns int64  `yaml:"maxIdleConns"`
	MaxOpenConns int64  `yaml:"maxOpenConns"`
}

type JwtConfig struct {
	JwtSecret string `yaml:"jwtSecret"`
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
