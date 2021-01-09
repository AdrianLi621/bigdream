package service

import "bigdream/huigou/model"

func SelectSeckill(condition map[string]interface{}, page int, pageSize int, orderBy string) []model.Seckill {
	return model.SelectSeckill(condition, page, pageSize, orderBy)
}
func CountSeckill(condition map[string]interface{}) int64 {
	return model.CountSeckill(condition)
}

