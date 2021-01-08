package controller

import (
	"bigdream/huigou/app/api/service"
	"bigdream/huigou/initialize"
	"bigdream/huigou/model"
	. "bigdream/huigou/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SelectStore struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`
}
type GetStoreGoods struct {
	StoreID     int `form:"store_id" json:"store_id"  binding:"required"`
}
type GetGoodsInfo struct {
	CommonId     int `form:"common_id" json:"common_id"  binding:"required"`
}
/**
获取店铺列表
*/
func StoreList(ctx *gin.Context) {

	var Store SelectStore
	if err := ctx.ShouldBind(&Store); err != nil {
		BadResponse(ctx, 0, nil, err.Error())
		return
	}
	condition := make(map[string]interface{})
	condition["store_state"] = 1
	count := service.CountStore(condition)
	var list []model.Store
	if count > 0 {
		list = service.SelectStore(condition, Store.Page, Store.PageSize, "store_sort desc")
	}
	response := make(map[string]interface{})
	response["page"] = Store.Page
	response["page_size"] = Store.PageSize
	response["list"] = list
	response["count"] = count

	SuccessResponse(ctx, 1, response, "获取成功")
}
/**
获取产品列表
*/
func StoreGoodsList(ctx *gin.Context) {
	var list GetStoreGoods
	if err := ctx.ShouldBindJSON(&list); err != nil {
		BadResponse(ctx, 0, nil, err.Error())
		return
	}
	condition := make(map[string]interface{})
	condition["store_id"] = list.StoreID
	condition["is_delete"] = 0
	var data []model.GoodsCommon
	count := service.CountGoodsCommon(condition)
	if count > 0 {
		data=service.SelectGoodsByCommonid(condition,0,0,"")
	}
	response := make(map[string]interface{})

	response["list"] = data
	response["count"] = count

	SuccessResponse(ctx, 1, response, "")

}
/**
获取产品详情
*/
func StoreGoodsInfo(ctx *gin.Context) {
	var list GetGoodsInfo
	if err := ctx.ShouldBindJSON(&list); err != nil {
		BadResponse(ctx, 0, nil, err.Error())
		return
	}
	condition := make(map[string]interface{})
	condition["common_id"] = list.CommonId
	condition["is_delete"] = 0
	var data []model.GoodsCommon
	count := service.CountGoodsCommon(condition)

	logger:=initialize.InitLogger("555555.log","info")
	logger.Info("打印",zap.Int64("line",count))

	if count > 0 {
		data=service.SelectGoodsByCommonid(condition,0,0,"")
	}
	response := make(map[string]interface{})

	response["list"] = data
	response["count"] = count

	SuccessResponse(ctx, 1, response, "")

}
