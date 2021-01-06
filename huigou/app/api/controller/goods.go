package controller

import (
	"bigdream/huigou/app/api/service"
	. "bigdream/huigou/pkg"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GoodsForm struct {
	GoodsName      string   `form:"goods_name" json:"goods_name" binding:"required"`
	GoodsBarcode   string   `form:"goods_barcode" json:"goods_barcode" binding:"required"`
	GoodsGcId      int      `form:"goods_gc_id" json:"goods_gc_id" binding:"required"`
	GoodsGcId1     int      `form:"goods_gc_id1" json:"goods_gc_id1" binding:"required"`
	GoodsGcId2     int      `form:"goods_gc_id2" json:"goods_gc_id2" binding:"required"`
	GoodsGcId3     int      `form:"goods_gc_id3" json:"goods_gc_id3" binding:"required"`
	GoodsGcName    string   `form:"goods_gc_name" json:"goods_gc_name" binding:"required"`
	GoodsPrice     int      `form:"goods_price" json:"goods_price" binding:"required"`
	GoodsInventory int      `form:"goods_inventory" json:"goods_inventory" binding:"required"`
	GoodsState     int      `form:"goods_state" json:"goods_state" `
	Attrs          int      `form:"attr_id" json:"attr_id"`
	AttrValueId    int      `form:"attr_value_id" json:"attr_value_id"`
	AttrValueName  string   `form:"attr_value_name" json:"attr_value_name"`
	GoodsImage     []string `form:"goods_image" json:"goods_image"`
	GoodsDescribe  string   `form:"goods_describe" json:"goods_describe" binding:"required"`
	GoodsImgUrl    string   `form:"goods_img_url" json:"goods_img_url" binding:"required"`
	GoodsImgSort   int      `form:"goods_img_sort" json:"goods_img_sort"`
}

func CreateGoods(ctx *gin.Context) {
	var goods GoodsForm
	if err := ctx.ShouldBind(&goods); err != nil {
		BadResponse(ctx, 0, nil, err.Error())
		return
	}

	common_data := make(map[string]interface{})
	common_data["goods_name"] = goods.GoodsName
	common_data["goods_gc_id"] = goods.GoodsGcId
	common_data["goods_gc_id1"] = goods.GoodsGcId1
	common_data["goods_gc_id2"] = goods.GoodsGcId2
	common_data["goods_gc_id3"] = goods.GoodsGcId3
	common_data["goods_image"] = goods.GoodsImgUrl
	common_data["goods_describe"] = goods.GoodsDescribe
	common_data["goods_gc_name"] = goods.GoodsGcName

	//fmt.Println(common_data)
	common_id := service.InsertGoodsCommon(common_data)

	fmt.Println(common_id)

	goods_data := make(map[string]interface{})
	goods_data["goods_commonid"] = common_id
	goods_data["goods_name"] = goods.GoodsName
	goods_data["goods_gc_id"] = goods.GoodsGcId
	goods_data["goods_gc_id1"] = goods.GoodsGcId1
	goods_data["goods_gc_id2"] = goods.GoodsGcId2
	goods_data["goods_gc_id3"] = goods.GoodsGcId3
	goods_data["goods_gc_name"] = goods.GoodsGcName
	goods_data["goods_sku"] = goods.GoodsGcName
	goods_data["goods_barcode"] = goods.GoodsGcName
	goods_data["goods_price"] = goods.GoodsPrice
	goods_data["goods_inventory"] = goods.GoodsInventory

	goods_data["attr_id"] = goods.GoodsInventory
	goods_data["attr_name"] = goods.GoodsInventory
	goods_data["attr_value_id"] = goods.GoodsInventory
	goods_data["attr_value_name"] = goods.GoodsInventory

	service.InsertGoods(goods_data)

}
