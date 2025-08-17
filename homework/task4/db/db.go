package db

import (
	"github.com/gin-gonic/gin"
	"github.com/go-master/task4/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var DB *gorm.DB

func InitDB(config model.MysqlConfig) error {
	var err error
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	//连接数据库
	DB, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().In(loc)
		},
	})
	if err != nil {
		return err
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(int(config.MaxOpenConns))
	sqlDB.SetMaxIdleConns(int(config.MaxIdleConns))
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	//创建表
	err = DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		return err
	}
	return err
}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 || pageSize > 100 {
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
