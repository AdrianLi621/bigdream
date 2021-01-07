package router

import (
	"bigdream/huigou/app/api/controller"

	"github.com/gin-gonic/gin"
)

/**
api路由配置
*/
func ApiRouter(group *gin.RouterGroup) {
	r := group.Group("")
	{
		r.POST("get_store_list", controller.StoreList)
		r.POST("add_goods", controller.CreateGoods)
		r.POST("get_store_goods", controller.StoreGoodsList)
		r.POST("get_goods_info", controller.StoreGoodsInfo)
	}
}
