package router

import (
	"bigdream/huigou/app/admin/controller"
	"github.com/gin-gonic/gin"
)

func AdminRouter(group *gin.RouterGroup) {
	r := group.Group("")
	{
		r.POST("/get_goods_class", controller.GoodsClassList)
		r.POST("/add_goods_class", controller.AddGoodsClass)
	}
}
