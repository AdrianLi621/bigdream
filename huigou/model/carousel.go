package model

import (
	. "bigdream/huigou/initialize"
	"fmt"
	"time"
)

type Carousel struct {
	Id              int       `gorm:"id;AUTO_INCREMENT;primary_key" json:"id" `
	ImgUrl          string    `gorm:"img_url" json:"src"`
	BackgroundColor string    `gorm:"background_color" json:"background"`
	IsShow          int       `gorm:"is_show" json:"-"`
	IsDelete        int       `gorm:"is_delete" json:"-"`
	AddTime         time.Time `gorm:"add_time" json:"-"`
	UpdateTime      time.Time `gorm:"update_time" json:"-"`
}

/**
插入轮播图
*/
func InsertCarousel(data map[string]interface{}) int {
	var carousel = &Carousel{
		ImgUrl:          data["img_url"].(string),
		BackgroundColor: data["background_color"].(string),
		IsShow:          data["is_show"].(int),
		IsDelete:        0,
		AddTime:         time.Now(),
		UpdateTime:      time.Now(),
	}
	err := DB.Create(carousel)
	if err.Error != nil {
		fmt.Println("插入产品失败", err.Error)
	}
	return carousel.Id
}

/**
统计数量
*/
func CountCarousel(condition map[string]interface{}) int64 {
	var count int64
	query := DB.Model(&Carousel{}).Where("1=1")
	if _, ok := condition["is_show"]; ok {
		query = query.Where("is_show", condition["is_show"])
	}
	if _, ok := condition["is_delete"]; ok {
		query = query.Where("is_delete", condition["is_delete"])
	}
	query.Count(&count)
	return count
}

/**
查询所有轮播图
*/
func SelectCarousel(condition map[string]interface{}, page int, pageSize int, orderBy string) []Carousel {
	var sto []Carousel
	query := DB.Model(&Carousel{}).Where("1=1")
	if _, ok := condition["is_show"]; ok {
		query = query.Where("is_show", condition["is_show"])
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
