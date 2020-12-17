package router

import (
	"github.com/gin-gonic/gin"
)

/**
api路由配置
*/
func ApiRouter(group *gin.RouterGroup) {
	r := group.Group("")
	{
		r.GET("/", func(context *gin.Context) {

		})
	}
}
