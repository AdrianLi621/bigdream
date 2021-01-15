package main

import (
	"bigdream/huigou/app/api/service"
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

func main()  {
	go CusGoodsFromMq()
	go CusGoodsCommonFromMq()
	for{
		select{

		}
	}
}
/**
详细产品同步
 */
func CusGoodsFromMq(){
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Error("消费mq队列，创建通道失败")
		return
	}

	err = ch.ExchangeDeclare(
		"ex_shelve_goods_common", // name
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
		"shelve_goods_common",    // routing key
		"ex_shelve_goods_common", // exchange
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
	forever:=make(chan int)
	for i:=0;i<100;i++ {
		go func(i int) {
			for d := range msgs {
				where:=make(map[string]interface{})
				common_id,_:=strconv.Atoi(string(d.Body))
				where["common_id"]=common_id
				isOk,_:=service.MigrateGoodsToES(where)
				if !isOk {
					fmt.Println("同步产品失败",err)
					//此处加入失败队列
				}
			}
		}(i)
	}
	fmt.Println("产品正在执行...")
	<-forever

}
/**
产品资料同步
 */
func CusGoodsCommonFromMq(){
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Error("消费mq队列，创建通道失败")
		return
	}

	err = ch.ExchangeDeclare(
		"ex_shelve_goods_common", // name
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
		"shelve_goods_common",    // routing key
		"ex_shelve_goods_common", // exchange
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
	forever:=make(chan int)
	for i:=0;i<100;i++ {
		go func(i int) {
			for d := range msgs {
				where:=make(map[string]interface{})
				common_id,_:=strconv.Atoi(string(d.Body))
				where["common_id"]=common_id
				isOk,_:=service.MigrateGoodsCommonToES(where)
				if !isOk {
					fmt.Println("同步产品资料失败",err)
					//此处加入失败队列
				}
			}
		}(i)
	}
	fmt.Println("产品资料正在执行...")
	<-forever

}