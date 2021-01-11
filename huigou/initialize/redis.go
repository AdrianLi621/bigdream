package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var ctx = context.Background()
var client *redis.Client
var sync_once sync.Once

//初始化
func init() {
	InitRedis()
}

/**
单例redis连接
*/
func InitRedis() *redis.Client {
	if client == nil {
		sync_once.Do(func() {
			client = redis.NewClient(&redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			})
		})
	}
	return client
}

//设置值
func SetRedisValue(key string, value interface{}) error {
	return client.Set(ctx, key, value, 0).Err()
}
//删除值
func DelRedisValue(key string) error {
	return client.Del(ctx, key).Err()
}
//获取值
func GetRedisValue(key string) (string, error) {
	return client.Get(ctx, key).Result()
}

//设置过期时间(单位 秒)
func SetRedisEXValue(key string, value interface{}, seconds time.Duration) (string, error) {
	return client.SetEX(ctx, key, value, seconds).Result()
}


//设置自增
func SetRedisIncr(key string) (int64, error) {
	return client.Incr(ctx,key).Result()
}


//设置值，类似加锁
func SetRedisNXValue(key string, value interface{}, seconds time.Duration) (bool, error) {
	return client.SetNX(ctx,key,value,seconds).Result()
}


