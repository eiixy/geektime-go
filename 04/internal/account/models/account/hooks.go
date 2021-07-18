package account

import "gorm.io/gorm"

func (account *Account) AfterDelete(tx *gorm.DB) (err error) {
	tx.Delete(&AccountInfo{}, account.ID)
	return
}
