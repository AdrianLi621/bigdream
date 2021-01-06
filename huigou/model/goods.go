package model

import (
	. "bigdream/huigou/initialize"
	"fmt"
	"time"
)

/**

 */
type Goods struct {
	GoodsId        int       `gorm:"goods_id;AUTO_INCREMENT" json:"goods_id" `
	GoodsCommonid  int       `gorm:"goods_commonid" json:"goods_commonid"`
	GoodsSku       string    `grom:"goods_sku" json:"goods_sku"`
	GoodsName      string    `grom:"goods_name" json:"goods_name"`
	GoodsBarcode   string    `gorm:"goods_barcode" json:"goods_barcode"`
	GoodsGcId      int       `grom:"goods_gc_id" json:"goods_gc_id"`
	GoodsGcId1     int       `gorm:"goods_gc_id1" json:"goods_gc_id1"`
	GoodsGcId2     int       `gorm:"goods_gc_id2" json:"goods_gc_id2"`
	GoodsGcId3     int       `gorm:"goods_gc_id3" json:"goods_gc_id3"`
	GoodsGcName    string    `gorm:"goods_gc_name" json:"goods_gc_name"`
	GoodsPrice     int       `grom:"goods_price" json:"goods_price"`
	GoodsSalenum   int       `grom:"goods_salenum" json:"goods_salenum"`
	GoodsInventory int       `gorm:"goods_inventory" json:"goods_inventory"`
	GoodsState     int       `grom:"goods_state" json:"goods_state"`
	AttrId         int       `gorm:"attr_id" json:"attr_id"`
	AttrName       string    `gorm:"attr_name" json:"attr_name"`
	AttrValueId    int       `gorm:"attr_value_id" json:"attr_value_id"`
	AttrValueName  string    `gorm:"attr_value_name" json:"attr_value_name"`
	IsDelete       int       `grom:"is_delete" json:"is_delete"`
	AddTime        time.Time `grom:"add_time" json:"add_time"`
	UpdateTime     time.Time `gorm:"update_time" json:"update_time"`
}

/**
插入产品
*/
func InsertGoods(data map[string]interface{}) int {
	var goods = &Goods{
		GoodsCommonid:  data["goods_commonid"].(int),
		GoodsSku:       data["goods_sku"].(string),
		GoodsName:      data["goods_name"].(string),
		GoodsBarcode:   data["goods_barcode"].(string),
		GoodsGcId:      data["goods_gc_id"].(int),
		GoodsGcId1:     data["goods_gc_id1"].(int),
		GoodsGcId2:     data["goods_gc_id2"].(int),
		GoodsGcId3:     data["goods_gc_id3"].(int),
		GoodsGcName:    data["goods_gc_name"].(string),
		GoodsPrice:     data["goods_price"].(int),
		GoodsSalenum:   0,
		GoodsInventory: data["goods_inventory"].(int),
		GoodsState:     1,
		AttrId:         data["attr_id"].(int),
		AttrName:       data["attr_name"].(string),
		AttrValueId:    data["attr_value_id"].(int),
		AttrValueName:  data["attr_value_name"].(string),
		IsDelete:       0,
		AddTime:        time.Now(),
		UpdateTime:     time.Now(),
	}
	err := DB.Create(goods)
	if err != nil {
		fmt.Println(err)
		fmt.Println("插入产品失败")
	}
	return goods.GoodsId
}
