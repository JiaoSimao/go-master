package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	var err error
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
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
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return err
}
