package controller

import (
	"bigdream/huigou/app/api/service"
	"bigdream/huigou/model"
	. "bigdream/huigou/pkg"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type GoodsForm struct {
	GoodsName     string `form:"goods_name" json:"goods_name" binding:"required"`
	StoreId       int    `form:"store_id" json:"store_id" binding:"required"`
	GoodsGcId     int    `form:"goods_gc_id" json:"goods_gc_id" binding:"required"`
	GoodsGcId1    int    `form:"goods_gc_id1" json:"goods_gc_id1" binding:"required"`
	GoodsGcId2    int    `form:"goods_gc_id2" json:"goods_gc_id2" binding:"required"`
	GoodsGcId3    int    `form:"goods_gc_id3" json:"goods_gc_id3" binding:"required"`
	GoodsGcName   string `form:"goods_gc_name" json:"goods_gc_name" binding:"required"`
	GoodsState    int    `form:"goods_state" json:"goods_state" `
	GoodsDescribe string `form:"goods_describe" json:"goods_describe" binding:"required"`
	GoodsImgUrl   string `form:"goods_img_url" json:"goods_img_url" binding:"required"`
	GoodsImgSort  int    `form:"goods_img_sort" json:"goods_img_sort"`
	GoodsInfo     []Info `json:"goods_info" form:"goods_info"`
	GoodsAttr     []Attr `json:"goods_attr" form:"goods_attr"`
}
type Attr struct {
	AttrId   int    `form:"attr_id" json:"attr_id"`
	AttrName string `form:"attr_name" json:"attr_name"`
	Child    []struct {
		AttrValueId   int    `form:"attr_value_id" json:"attr_value_id"`
		AttrValueName string `form:"attr_value_name" json:"attr_value_name"`
	}
}

type Info struct {
	GoodsPrice     int      `form:"goods_price" json:"goods_price"`
	GoodsInventory int      `form:"goods_inventory" json:"goods_inventory"`
	GoodsImages    []string `form:"goods_images" json:"goods_images"`
	Spec           []struct {
		AttrValueId   int    `form:"attr_value_id" json:"attr_value_id"`
		AttrValueName string `form:"attr_value_name" json:"attr_value_name"`
	}
}

/**
创建产品
*/
func CreateGoods(ctx *gin.Context) {
	var goods GoodsForm
	if err := ctx.ShouldBindJSON(&goods); err != nil {
		BadResponse(ctx, 0, nil, err.Error())
		return
	}
	bytes, err := json.Marshal(goods.GoodsAttr)
	if err != nil {
		BadResponse(ctx, 0, nil, err.Error())
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
	common_data["store_id"] = goods.StoreId
	common_data["goods_spec"] = string(bytes)

	common_id := service.InsertGoodsCommon(common_data)
	snow_flake:=NewSnowFlake()
	
	var goods_ids []int

	if len(goods.GoodsInfo) > 0 {
		for _, v := range goods.GoodsInfo {
			goods_sku:=snow_flake.MakeUniqueId()
			goods_barcode:=snow_flake.MakeUniqueId()
			spec, err := json.Marshal(v.Spec)
			if err != nil {
				BadResponse(ctx, 0, nil, err.Error())
			}
			var str_name string
			for _, v := range v.Spec {
				str_name += " " + v.AttrValueName
			}
			goods_data := make(map[string]interface{})
			goods_data["goods_commonid"] = common_id
			goods_data["goods_sku"] = goods_sku
			goods_data["goods_barcode"] = goods_barcode
			goods_data["goods_name"] = goods.GoodsName + str_name
			goods_data["goods_gc_id"] = goods.GoodsGcId
			goods_data["goods_gc_id1"] = goods.GoodsGcId1
			goods_data["goods_gc_id2"] = goods.GoodsGcId2
			goods_data["goods_gc_id3"] = goods.GoodsGcId3
			goods_data["goods_gc_name"] = goods.GoodsGcName
			goods_data["goods_price"] = v.GoodsPrice
			goods_data["goods_inventory"] = v.GoodsInventory
			goods_data["goods_spec"] = string(spec)
			goods_data["store_id"] = goods.StoreId
			goods_id := service.InsertGoods(goods_data)
			goods_ids=append(goods_ids,goods_id)
			for _, val := range v.GoodsImages {
				image_data := make(map[string]interface{})
				image_data["goods_id"] = goods_id
				image_data["goods_img_url"] = val
				service.InsertGoodsImages(image_data)
			}
		}
	}
	if len(goods_ids)>0 {
		where,row:=make(map[string]interface{}),make(map[string]interface{})
		where["common_id"]=common_id
		row["goods_id"]=goods_ids[0]
		service.UpGoodsCommon(where,row)
	}
	SuccessResponse(ctx, 0, nil, "添加成功")
}
/**
搜索
 */
func Search(ctx *gin.Context)  {
	a,_:=SelectDoc("student")
	var t model.GoodsCommon
	goods_data := make(map[string]interface{})
	var num int
	for _,v:=range a{
		err:=json.Unmarshal(v.Source,&t)
		if err != nil {
			BadResponse(ctx, 0, nil, err.Error())
		}
		goods_data["goods_name"] = t.GoodsName
		goods_data["goods_image"] = t.GoodsImage
		goods_data["goods_gc_name"] = t.GoodsGcName
		num++
	}
	response := make(map[string]interface{})
	response["list"] = goods_data
	response["count"] = num
	SuccessResponse(ctx, 0, response, "获取成功")
}