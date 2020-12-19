package service

import "bigdream/huigou/model"

/**
插入菜单分类
*/
func InsertGoodsClass(data map[string]interface{}) int {
	return model.InsertGoodsClass(data)
}

func SelectGoodsClass(condition map[string]interface{}, page int, pageSize int, orderBy string) []model.GoodsClass {
	return model.SelectGoodsClass(condition, page, pageSize, orderBy)
}
func CountGoodsClass(condition map[string]interface{}) int64{
	return model.CountGoodsClass(condition)
}