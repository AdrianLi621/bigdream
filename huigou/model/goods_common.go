package model

import (
	. "bigdream/huigou/initialize"
	"bigdream/huigou/pkg"
	"encoding/json"
	"fmt"
	"time"
)

/**

 */
type GoodsCommon struct {
	GoodsCommonid int       `gorm:"goods_commonid;AUTO_INCREMENT;primary_key" json:"goods_commonid" `
	GoodsName     string    `gorm:"goods_name" json:"goods_name"`
	StoreId       int       `grom:"store_id" json:"store_id"`
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
	GoodsList  []Goods      `gorm:"FOREIGNKEY:goods_commonid" json:"goods_list"`
}
type EsGoodsCommon struct {
	GoodsCommonid int       `json:"goods_commonid" `
	GoodsName     string    `json:"goods_name"`
	StoreId       int       `json:"store_id"`
	GoodsGcId     int       `json:"goods_gc_id"`
	GoodsGcId1    int       `json:"goods_gc_id1"`
	GoodsGcId2    int       `json:"goods_gc_id2"`
	GoodsGcId3    int       `json:"goods_gc_id3"`
	GoodsGcName   string    `json:"goods_gc_name"`
	GoodsImage    string    `json:"goods_image"`
	GoodsState    int       `json:"goods_state"`
	IsDelete      int       `json:"is_delete"`
	AddTime       time.Time `json:"add_time"`
	UpdateTime    time.Time `json:"update_time"`
}

/**
插入产品
*/
func InsertGoodsCommon(data map[string]interface{}) int {
	var goods_common = &GoodsCommon{
		GoodsName:     data["goods_name"].(string),
		StoreId:       data["store_id"].(int),
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
/**
统计数量
*/
func CountGoodsCommon(condition map[string]interface{}) int64 {
	var count int64
	query := DB.Model(&GoodsCommon{}).Where("1=1")
	if _, ok := condition["store_id"]; ok {
		query = query.Where("store_id", condition["store_id"])
	}
	if _, ok := condition["is_delete"]; ok {
		query = query.Where("is_delete", condition["is_delete"])
	}
	if _, ok := condition["common_id"]; ok {
		query = query.Where("goods_commonid", condition["common_id"])
	}
	query.Count(&count)
	return count
}

/**
查询所有店铺
*/
func SelectGoodsCommon(condition map[string]interface{}, page int, pageSize int, orderBy string) []GoodsCommon {
	var sto []GoodsCommon
	query := DB.Model(&GoodsCommon{}).Where("1=1")
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
/**
查询所有店铺 by commonid
*/
func SelectGoodsByCommonid(condition map[string]interface{}, page int, pageSize int, orderBy string) []GoodsCommon {
	var sto []GoodsCommon
	query := DB.Model(&GoodsCommon{}).Preload("GoodsList").Where("1=1")
	if _, ok := condition["store_id"]; ok {
		query = query.Where("store_id", condition["store_id"])
	}
	if _, ok := condition["is_delete"]; ok {
		query = query.Where("is_delete", condition["is_delete"])
	}
	if _, ok := condition["common_id"]; ok {
		query = query.Where("goods_commonid", condition["common_id"])
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
/**
更新产品
*/
func UpGoodsCommon(condition map[string]interface{},row map[string]interface{}){
	query := DB.Model(&GoodsCommon{})
	if _, ok := condition["common_id"]; ok {
		query = query.Where("goods_commonid", condition["common_id"])
	}
	query.Updates(row)
}
/**
把数据抛入es
*/
func MigrateGoodsCommonToES(condition map[string]interface{})(bool,error)  {
	var sto EsGoodsCommon
	query := DB.Model(&GoodsCommon{}).Where("1=1")
	if _, ok := condition["common_id"]; ok {
		query = query.Where("goods_commonid", condition["common_id"])
	}
	query.Find(&sto)
	bytes,_:=json.Marshal(sto)
	_,err:=pkg.InsertDoc("goods_common",string(bytes))
	if err != nil{
		return false,err
	}
	return true,nil
}