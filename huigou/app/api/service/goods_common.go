package service

import "bigdream/huigou/model"

func InsertGoodsCommon(data map[string]interface{}) int {
	return model.InsertGoodsCommon(data)
}
func SelectGoodsCommon(condition map[string]interface{}, page int, pageSize int, orderBy string) []model.GoodsCommon {
	return model.SelectGoodsCommon(condition, page, pageSize, orderBy)
}
func CountGoodsCommon(condition map[string]interface{}) int64 {
	return model.CountGoodsCommon(condition)
}
func SelectGoodsByCommonid(condition map[string]interface{}, page int, pageSize int, orderBy string) []model.GoodsCommon {
	return model.SelectGoodsByCommonid(condition, page, pageSize, orderBy)
}
func UpGoodsCommon(where map[string]interface{},row map[string]interface{})  {
	model.UpGoodsCommon(where,row)
}