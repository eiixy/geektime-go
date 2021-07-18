package account

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email    string `gorm:"comment:邮箱号"`
	Password string `gorm:"comment:密码"`
}
