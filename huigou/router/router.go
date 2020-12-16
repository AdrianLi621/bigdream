package router

import (
	api_router "bigdream/huigou/app/api/router"
	"github.com/gin-gonic/gin"
)

/**
统一路由配置
*/
func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	api := router.Group("api")
	{
		api_router.ApiRouter(api)
	}
	router.Run(":8888")
	return router
}
