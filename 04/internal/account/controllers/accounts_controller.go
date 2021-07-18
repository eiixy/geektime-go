package controllers

import (
	"geektime-go/04/internal/account/models/account"
	"geektime-go/04/internal/account/models/scopes"
	"geektime-go/04/internal/account/services"
	"geektime-go/04/pkg/database"
	"github.com/gin-gonic/gin"
)

func GetAccounts(context *gin.Context) {
	accounts := database.DB.Scopes(scopes.Paginate(context)).Select(&account.Account{})
	success(context, accounts)
}
func CreateAccount(context *gin.Context) {
	a := &account.Account{
		Email: context.PostForm("email"),
	}
	services.Register(*a)
	success(context, nil)
}
