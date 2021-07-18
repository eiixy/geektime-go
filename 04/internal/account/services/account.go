package services

import (
	"geektime-go/04/internal/account/models/account"
	"geektime-go/04/pkg/database"
)

func Register(account account.Account) {
	database.DB.Create(account)
}
