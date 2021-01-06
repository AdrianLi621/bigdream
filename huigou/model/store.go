package model

import (
	. "bigdream/huigou/initialize"
	"time"
)

/**
店铺
*/
type Store struct {
	StoreId        int       `gorm:"store_id;AUTO_INCREMENT" json:"store_id" `
	StoreName      string    `gorm:"store_name" json:"store_name"`
	StoreLogo      string    `grom:"store_logo" json:"store_logo"`
	ProvinceId     int       `grom:"province_id" json:"province_id"`
	CityId         int       `grom:"city_id" json:"city_id"`
	AreaId         int       `gorm:"area_id" json:"area_id"`
	AreaDetail     string    `grom:"area_detail" json:"area_detail"`
	AddressInfo    string    `gorm:"address_info" json:"address_info"`
	StoreState     int       `gorm:"store_state" json:"store_state"`
	StoreSort      int       `gorm:"store_sort" json:"store_sort"`
	StoreOpenTime  time.Time `gorm:"store_open_time" json:"store_open_time"`
	StoreCloseTime time.Time `gorm:"store_close_time" json:"store_close_time"`
	StoreKeywords  time.Time `gorm:"store_keywords" json:"store_keywords"`
	StorePhone     string    `gorm:"store_phone" json:"store_phone"`
	AddTime        time.Time `gorm:"add_time" json:"add_time"`
	UpdateTime     time.Time `gorm:"update_time" json:"update_time"`
}

/**
统计数量
*/
func CountStore(condition map[string]interface{}) int64 {
	var count int64
	query := DB.Model(&Store{}).Where("1=1")
	// if _, ok := condition["is_delete"]; ok {
	// 	query = query.Where("is_delete", condition["is_delete"])
	// }
	// if _, ok := condition["gc_name_like"]; ok {
	// 	query = query.Where("gc_name like ? ", condition["gc_name_like"])
	// }
	query.Count(&count)
	return count
}

/**
查询所有店铺
*/
func SelectStore(condition map[string]interface{}, page int, pageSize int, orderBy string) []Store {
	var sto []Store
	query := DB.Model(&Store{}).Where("1=1")
	if _, ok := condition["store_state"]; ok {
		query = query.Where("store_state", condition["store_state"])
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
