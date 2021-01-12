package router

import (
	"bigdream/huigou/app/api/controller"
	"bigdream/huigou/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
api路由配置
*/

func ApiRouter(group *gin.RouterGroup) {
	r := group.Group("")
	{
		r.POST("get_store_list", controller.StoreList)       //获取店铺列表
		r.POST("add_goods", controller.CreateGoods)          //添加产品
		r.POST("get_store_goods", controller.StoreGoodsList) //获取店铺产品
		r.POST("get_goods_info", controller.StoreGoodsInfo)  // 产品详情
		r.POST("get_carousel_list", controller.CarouselList) //轮播图列表
		r.POST("get_seckill_list", controller.SeckillList) //秒杀列表
		r.Any("search_goods", controller.Search)//搜搜产品
		r.Any("test", func(ctx *gin.Context) {
			for i:=0;i<30;i++{
				res,err:=pkg.ProToWorks("rereir",strconv.Itoa(i))
				if err !=nil {
					panic(err)
				}
				fmt.Println(res)
			}

		})
	}
}
