package scopes

import "gorm.io/gorm"

func SoftDeletes(db *gorm.DB) *gorm.DB {
	return db.Where("delete_at IS NULL")
}
