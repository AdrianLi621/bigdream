package controller



import (
	"bigdream/huigou/app/api/service"
	"bigdream/huigou/model"
	. "bigdream/huigou/pkg"

	"github.com/gin-gonic/gin"
)




/**
获取轮播图列表
*/
func CarouselList(ctx *gin.Context) {
	condition := make(map[string]interface{})
	condition["is_show"] = 1
	condition["is_delete"] = 0
	var data []model.Carousel
	count := service.CountCarousel(condition)
	if count > 0 {
		data=service.SelectCarousel(condition,0,0,"")
	}
	response := make(map[string]interface{})
	response["list"] = data
	response["count"] = count
	SuccessResponse(ctx, 1, response, "")

}
























