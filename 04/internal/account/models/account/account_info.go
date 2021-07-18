package account

import "geektime-go/04/internal/account/models/relations"

type AccountInfo struct {
	AccountId uint   `gorm:"uniqueIndex"`
	NikeName  string `gorm:"comment:昵称"`
	relations.HasAccount
}
