package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Dao dao
type Dao struct {
	db *gorm.DB
}

// New init mysql db
func New() (dao *Dao) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/geektime-go?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&User{})
	return &Dao{db}
}
