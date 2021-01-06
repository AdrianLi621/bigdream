package controller

import (
	"bigdream/huigou/app/api/service"
	"bigdream/huigou/model"
	. "bigdream/huigou/pkg"

	"github.com/gin-gonic/gin"
)

type SelectStore struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`
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

	SuccessResponse(ctx, 1, response, "")
}
