package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	code    int
	message string
	data    interface{}
}

func json(ctx *gin.Context, response response) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    response.code,
		"message": response.message,
		"data":    response.data,
	})
}

func success(ctx *gin.Context, data interface{}) {
	json(ctx, response{
		code:    0,
		message: "success",
		data:    data,
	})
}

func failed(ctx *gin.Context, message string, code int) {
	json(ctx, response{
		code:    code,
		message: message,
	})
}
