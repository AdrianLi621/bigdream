package model

import (
	. "bigdream/huigou/initialize"
	"time"
)

/**
店铺
*/
type Seckill struct {
	SecId         int       `gorm:"sec_id;AUTO_INCREMENT;primary_key" json:"sec_id" `
	SecTitle      string    `gorm:"sec_title" json:"sec_title"`
	GoodsId       int       `gorm:"goods_id" json:"goods_id"`
	GoodsCommonid int       `gorm:"goods_commonid" json:"goods_commonid"`
	StartTime     time.Time `gorm:"start_time" json:"start_time"`
	EndTime       time.Time `gorm:"end_time" json:"end_time"`
	SecPrice      int       `gorm:"sec_price" json:"sec_price"`
	SecInventory  int       `gorm:"sec_inventory" json:"sec_inventory"`
	IsDelete      int       `gorm:"is_delete" json:"is_delete"`
	State         int       `gorm:"state" json:"state"`
	AddTime       time.Time `gorm:"add_time" json:"add_time"`
	UpdateTime    time.Time `gorm:"update_time" json:"update_time"`
}

/**
统计数量
*/
func CountSeckill(condition map[string]interface{}) int64 {
	var count int64
	query := DB.Model(&Seckill{}).Where("1=1")
	if _, ok := condition["is_delete"]; ok {
		query = query.Where("is_delete", condition["is_delete"])
	}
	if _, ok := condition["state"]; ok {
		query = query.Where("state", condition["state"])
	}
	query.Count(&count)
	return count
}

/**
查询所有数据
*/
func SelectSeckill(condition map[string]interface{}, page int, pageSize int, orderBy string) []Seckill {
	var sto []Seckill
	query := DB.Model(&Seckill{}).Where("1=1")
	if _, ok := condition["is_delete"]; ok {
		query = query.Where("is_delete", condition["is_delete"])
	}
	if _, ok := condition["state"]; ok {
		query = query.Where("state", condition["state"])
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
