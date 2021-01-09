package pkg

import (
	"bigdream/huigou/initialize"
	"sync"
	"time"
)

var m sync.Mutex

type UniqueId interface {
	MakeUniqueId() string
}
//redis生成唯一id
type RedisEngine struct {
}

func NewRedisEngine() *RedisEngine {
	return &RedisEngine{}
}

func (r *RedisEngine) MakeUniqueId()string {
	res,_:=initialize.SetRedisNXValue("lock",1,time.Second)
	if res {
		now:=time.Now()
		str:=now.Format("20060102150405")
		return str
	}
	return ""
}
//雪花算法生成唯一id
type SnowFlake struct {
}

func NewSnowFlake() *SnowFlake {
	return &SnowFlake{}
}

func (s *SnowFlake) MakeUniqueId() {

}



//go语言加锁生成唯一id
type SystemEngine struct {
}

func NewSystemEngine() *SystemEngine {
	return &SystemEngine{}
}


func (r *SystemEngine) MakeUniqueId()string {
	m.Lock()
	defer m.Unlock()
	now:=time.Now()
	str:=now.Format("20060102150405")
	return str
}
