package pkg

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

var m sync.Mutex

type UniqueId interface {
	MakeUniqueId() string
}
////redis生成唯一id
//type RedisEngine struct {
//}
//
//func NewRedisEngine() *RedisEngine {
//	return &RedisEngine{}
//}
//
//func (r *RedisEngine) MakeUniqueId()string {
//	res,_:=initialize.SetRedisNXValue("lock",1,time.Second)
//	if res {
//		now:=time.Now()
//		str:=now.Format("20060102150405")
//		incr,err:=initialize.SetRedisIncr("incr")
//		if err != nil {
//			panic(err)
//		}
//		str=str+strconv.FormatInt(incr,10)
//		initialize.DelRedisValue("lock")
//		return str
//	}
//	return ""
//}
//雪花算法生成唯一id

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)
type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}
func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	return ID
}
type SnowFlake struct {

}

func NewSnowFlake() *SnowFlake {
	return &SnowFlake{}
}

func (s *SnowFlake) MakeUniqueId()string {
	node, err := NewWorker(1)
	if err != nil {
		panic(err)
	}
	return strconv.FormatInt(node.GetId(),10)
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
