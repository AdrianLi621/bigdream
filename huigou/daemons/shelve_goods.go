package main

import (
	"bigdream/huigou/initialize"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"strconv"
	"sync"
)

var conn *amqp.Connection
var err error
var once sync.Once
var log  *zap.Logger


func init() {
	if conn == nil {
		once.Do(func() {
			conn, err = amqp.Dial("amqp://guest:lijingqwer@121.5.22.154:5672/")
			if err != nil {
				panic(err)
			}
		})
	}
	if log == nil {
		once.Do(func() {
			log = initialize.InitLogger("CusGoodsFromMq.log", "error")
		})
	}

}
func main(){
	CusGoodsFromMq()
}
//守护进程    mq上架产品
/**
  路由模式  消费者
*/
func CusGoodsFromMq(){
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Error("消费mq队列，创建通道失败")
		return
	}

	err = ch.ExchangeDeclare(
		"ex_shelve_goods", // name
		"direct",          // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		log.Error("消费mq队列，创建交换机失败")
		return
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Error("消费mq队列，创建队列失败")
		return
	}

	err = ch.QueueBind(
		q.Name,            // queue name
		"shelve_goods",    // routing key
		"ex_shelve_goods", // exchange
		false,
		nil)
	if err != nil {
		log.Error("消费mq队列，队列绑定交换机失败")
		return
	}
	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error("消费mq队列，读取队列消息失败")
		return
	}
	forever := make(chan bool)
	for i:=0;i<100;i++ {
		go func(i int) {
			for d := range msgs {
				fmt.Println("协程"+strconv.Itoa(i)+"读取到消息:",string(d.Body))
			}
		}(i)
	}
	fmt.Println("等待读取消息......")
	<-forever

}
