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
		r.GET("get_store_list", controller.StoreList)
		r.POST("add_goods", controller.CreateGoods)
	}
}
