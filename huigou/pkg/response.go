package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Msg  string `json:"msg"`
}

func BadResponse(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(http.StatusBadRequest, &Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
	return
}
func SuccessResponse(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, &Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
	return
}
