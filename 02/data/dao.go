package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Dao dao
type Dao struct {
	db *sql.DB
}

// New init mysql db
func New() (dao *Dao) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/geektime-go?charset=utf8mb4"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql open error: %+v\r\n", err)
		panic("failed to connect database")
	}
	return &Dao{db: db}
}
