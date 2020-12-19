package model

import (
	. "bigdream/huigou/initialize"
	"fmt"
	"time"
)

/**
分类菜单
*/
type GoodsClass struct {
	GcId      int       `gorm:"gc_id;AUTO_INCREMENT" json:"gc_id" `
	GcName    string    `gorm:"gc_name;unique_index" json:"gc_name"`
	GcPid     int       `grom:"gc_pid" json:"gc_pid"`
	GcSort    int       `grom:"gc_sort" json:"gc_sort"`
	GcLevel   int       `gorm:"gc_level" json:"gc_level"`
	GcPids    string    `grom:"gc_pids" json:"gc_pids"`
	GcImgUrl  string    `gorm:"gc_img_url" json:"gc_img_url"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}
/**
统计数量
 */
func CountGoodsClass(condition map[string]interface{}) int64 {
	var count int64
	query := DB.Model(&GoodsClass{}).Where("1=1")
	if _, ok := condition["is_delete"]; ok {
		query = query.Where("is_delete", condition["is_delete"])
	}
	if _, ok := condition["gc_name_like"]; ok {
		query = query.Where("gc_name like ? ", condition["gc_name_like"])
	}
	query.Count(&count)
	return count
}

/**
查询所有分类
 */
func SelectGoodsClass(condition map[string]interface{}, page int, pageSize int, orderBy string) []GoodsClass {
	var gc []GoodsClass
	query := DB.Model(&GoodsClass{}).Where("1=1")
	if _, ok := condition["is_delete"]; ok {
		query = query.Where("is_delete", condition["is_delete"])
	}
	if _, ok := condition["gc_name_like"]; ok {
		query = query.Where("gc_name like ? ", condition["gc_name_like"])
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
	query.Find(&gc)
	return gc
}

/**
插入分类菜单
*/
func InsertGoodsClass(data map[string]interface{}) int {
	var gc = &GoodsClass{
		GcName:    data["gc_name"].(string),
		GcPid:     data["gc_pid"].(int),
		GcSort:    data["gc_sort"].(int),
		GcLevel:   data["gc_level"].(int),
		GcPids:    data["gc_pids"].(string),
		GcImgUrl:  data["gc_img_url"].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := DB.Create(gc)
	if err != nil {
		fmt.Println(err)
		fmt.Println("插入菜单分类失败")
	}
	return gc.GcId
}
