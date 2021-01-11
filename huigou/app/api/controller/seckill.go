package controller

import (
	"bigdream/huigou/app/api/service"
	"bigdream/huigou/model"
	. "bigdream/huigou/pkg"
	"github.com/gin-gonic/gin"
)

/**
获取秒杀列表
*/
func SeckillList(ctx *gin.Context) {
	condition := make(map[string]interface{})
	condition["state"] = 2
	condition["is_delete"] = 0
	var data []model.Seckill
	count := service.CountSeckill(condition)
	if count > 0 {
		data = service.SelectSeckill(condition, 0, 0, "")
	}
	response := make(map[string]interface{})
	response["list"] = data
	response["count"] = count
	SuccessResponse(ctx, 1, response, "")
}
