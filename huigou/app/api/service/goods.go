package service

import "bigdream/huigou/model"

func InsertGoods(data map[string]interface{}) int {
	return model.InsertGoods(data)
}
func SelectGoods(condition map[string]interface{}, page int, pageSize int, orderBy string) []model.Goods {
	return model.SelectGoods(condition, page, pageSize, orderBy)
}
func CountGoods(condition map[string]interface{}) int64 {
	return model.CountGoods(condition)
}
