package service


import "bigdream/huigou/model"


func SelectCarousel(condition map[string]interface{}, page int, pageSize int, orderBy string) []model.Carousel {
	return model.SelectCarousel(condition, page, pageSize, orderBy)
}
func CountCarousel(condition map[string]interface{}) int64 {
	return model.CountCarousel(condition)
}






















