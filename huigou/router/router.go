package router

import (
	admin_router "bigdream/huigou/app/admin/router"
	api_router "bigdream/huigou/app/api/router"
	"bigdream/huigou/middleware"
	"github.com/gin-gonic/gin"
)

/**
统一路由配置
*/
func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middleware.Cors())
	api := router.Group("api")
	{
		api_router.ApiRouter(api)
	}
	admin := router.Group("admin")
	{
		admin_router.AdminRouter(admin)
	}
	router.Run(":8888")
	return router
}
