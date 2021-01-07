package model

import (
	. "bigdream/huigou/initialize"
	"fmt"
	"time"
)

/**

 */
type GoodsCommon struct {
	GoodsCommonid int       `gorm:"goods_commonid;AUTO_INCREMENT;primary_key" json:"goods_commonid" `
	GoodsName     string    `gorm:"goods_name" json:"goods_name"`
	GoodsGcId     int       `grom:"goods_gc_id" json:"goods_gc_id"`
	GoodsGcId1    int       `grom:"goods_gc_id1" json:"goods_gc_id1"`
	GoodsGcId2    int       `gorm:"goods_gc_id2" json:"goods_gc_id2"`
	GoodsGcId3    int       `gorm:"goods_gc_id3" json:"goods_gc_id3"`
	GoodsGcName   string    `grom:"goods_gc_name" json:"goods_gc_name"`
	GoodsImage    string    `gorm:"goods_image" json:"goods_image"`
	GoodsSpec    string    `gorm:"goods_spec" json:"goods_spec"`
	GoodsDescribe string    `grom:"goods_describe" json:"goods_describe"`
	GoodsState    int       `grom:"goods_state" json:"goods_state"`
	IsDelete      int       `gorm:"is_delete" json:"is_delete"`
	AddTime       time.Time `grom:"add_time" json:"add_time"`
	UpdateTime    time.Time `gorm:"update_time" json:"update_time"`
}

/**
插入产品
*/
func InsertGoodsCommon(data map[string]interface{}) int {
	var goods_common = &GoodsCommon{
		GoodsName:     data["goods_name"].(string),
		GoodsGcId:     data["goods_gc_id"].(int),
		GoodsGcId1:    data["goods_gc_id1"].(int),
		GoodsGcId2:    data["goods_gc_id2"].(int),
		GoodsGcId3:    data["goods_gc_id3"].(int),
		GoodsGcName:   data["goods_gc_name"].(string),
		GoodsImage:    data["goods_image"].(string),
		GoodsDescribe: data["goods_describe"].(string),
		GoodsSpec: data["goods_spec"].(string),
		GoodsState:    1,
		IsDelete:      0,
		AddTime:       time.Now(),
		UpdateTime:    time.Now(),
	}
	err := DB.Create(goods_common)
	if err.Error != nil {
		fmt.Println("插入公共产品失败", err.Error)
	}
	return goods_common.GoodsCommonid
}
