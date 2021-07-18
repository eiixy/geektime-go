package database

import (
	"gorm.io/gorm"
)

// DB gorm.DB 对象
var DB *gorm.DB

func Connect(name ...string) *gorm.DB {
	if len(name) == 0 {
		DB = ConnectMysql()
	} else {
		switch name[0] {
		case "mysql":
			DB = ConnectMysql()
		default:
			DB = ConnectMysql()
		}
	}
	return DB
}
