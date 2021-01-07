package model

import (
	. "bigdream/huigou/initialize"
	"fmt"
	"time"
)

/**

 */
type Goods struct {
	GoodsId        int       `gorm:"goods_id;AUTO_INCREMENT;primary_key" json:"goods_id" `
	GoodsCommonid  int       `gorm:"goods_commonid" json:"goods_commonid"`
	StoreId     int    `form:"store_id" json:"store_id" binding:"required"`
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
	GoodsSpec      string     `gorm:"goods_spec" json:"goods_spec"`
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
		StoreId:        data["store_id"].(int),
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
		GoodsSpec: data["goods_spec"].(string),
		IsDelete:       0,
		AddTime:        time.Now(),
		UpdateTime:     time.Now(),
	}
	err := DB.Create(goods)
	if err.Error != nil {
		fmt.Println("插入产品失败",err.Error)
	}
	return goods.GoodsId
}
/**
统计数量
*/
func CountGoods(condition map[string]interface{}) int64 {
	var count int64
	query := DB.Model(&Goods{}).Where("1=1")
	if _, ok := condition["store_id"]; ok {
		query = query.Where("store_id", condition["store_id"])
	}
	if _, ok := condition["is_delete"]; ok {
		query = query.Where("is_delete", condition["is_delete"])
	}
	query.Count(&count)
	return count
}

/**
查询所有店铺
*/
func SelectGoods(condition map[string]interface{}, page int, pageSize int, orderBy string) []Goods {
	var sto []Goods
	query := DB.Model(&Goods{}).Where("1=1")
	if _, ok := condition["store_id"]; ok {
		query = query.Where("store_id", condition["store_id"])
	}
	if _, ok := condition["is_delete"]; ok {
		query = query.Where("is_delete", condition["is_delete"])
	}
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	query = query.Limit(pageSize).Offset((page - 1) * pageSize)
	if len(orderBy) > 0 {
		query = query.Order(orderBy)
	}
	query.Find(&sto)
	return sto
}
