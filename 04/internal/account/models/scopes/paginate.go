package scopes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func Paginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page == 0 {
		page = 1
	}

	size, _ := strconv.Atoi(ctx.Query("size"))
	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 10
	}
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
