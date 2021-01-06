package service

import "bigdream/huigou/model"

func InsertGoods(data map[string]interface{}) int {
	return model.InsertGoods(data)
}
