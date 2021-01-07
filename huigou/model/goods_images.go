package model

import (
	. "bigdream/huigou/initialize"
	"fmt"
	"time"
)

/**

 */
type GoodsImages struct {
	ImgId        int       `gorm:"img_id;AUTO_INCREMENT;primary_key" json:"img_id" `
	GoodsId      int       `gorm:"goods_id" json:"goods_id"`
	GoodsImgUrl  string    `grom:"goods_img_url" json:"goods_img_url"`
	GoodsImgSort int       `grom:"goods_img_sort" json:"goods_img_sort"`
	IsDelete     int       `grom:"is_delete" json:"is_delete"`
	AddTime      time.Time `grom:"add_time" json:"add_time"`
	UpdateTime   time.Time `gorm:"update_time" json:"update_time"`
}

/**

 */
func InsertGoodsImages(data map[string]interface{}) int {
	var goods_images = &GoodsImages{
		GoodsId:      data["goods_id"].(int),
		GoodsImgUrl:  data["goods_img_url"].(string),
		GoodsImgSort: 0,
		IsDelete:     0,
		AddTime:      time.Now(),
		UpdateTime:   time.Now(),
	}
	err := DB.Create(goods_images)
	if err.Error != nil {
		fmt.Println("插入产品图片失败",err.Error)
	}
	return goods_images.ImgId
}
