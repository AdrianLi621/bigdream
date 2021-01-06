package service

import "bigdream/huigou/model"

func SelectStore(condition map[string]interface{}, page int, pageSize int, orderBy string) []model.Store {
	return model.SelectStore(condition, page, pageSize, orderBy)
}
func CountStore(condition map[string]interface{}) int64 {
	return model.CountStore(condition)
}
