package pkg

import (
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
)

var Conn *amqp.Connection

func init() {
	if Conn == nil {
		Conn, err = amqp.Dial("amqp://guest:lijingqwer@121.5.22.154:5672/")
		if err != nil {
			panic(err)
		}
	}
}
/**
抛入待上架产品到mq
 */
func ShelveGoodsToMq(goods_id int)(bool,error)  {
	ch, err := Conn.Channel()
	defer ch.Close()
	if err != nil {
		return false, err
	}
	err=ch.ExchangeDeclare("ex_shelve_goods_common","direct",true,false,false,false,nil)
	if err != nil {
		return false, err
	}
	err=ch.Publish("ex_shelve_goods_common","shelve_goods_common",false,false,amqp.Publishing{
		Body: []byte(strconv.Itoa(goods_id)),
	})
	if err != nil {
		return false, err
	}
	return true,nil
}

















/**
1对多模式  生产者
*/
func ProToWorks(queue string, body string) (bool, error) {
	ch, err := Conn.Channel()
	defer ch.Close()
	if err != nil {
		return false, err
	}
	q,err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return false, err
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			Body: []byte(body),
		})
	if err != nil {
		return false, err
	}
	return true, nil
}
/**
消费者
 */
func CusToWorks(queue string) {
	defer Conn.Close()
	ch, err := Conn.Channel()
	if err != nil {
		panic(err)
	}
	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		panic(err)
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}

	for d := range msgs {//此处阻塞
		fmt.Println(string(d.Body))
		d.Ack(false)
	}

}
/**
交换机模式(广播模式)
 */
func ProToExc(exchange string, body string) (bool,error) {
	ch, err := Conn.Channel()
	defer ch.Close()
	if err != nil {
		return false, err
	}
	err = ch.ExchangeDeclare(
		exchange,   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return false, err
	}

	err = ch.Publish(
		exchange, // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			Body:        []byte(body),
		})
	if err != nil {
		return false, err
	}
	return true, nil

}
/**
交换机消费 （广播模式）
 */
func CusExc(exchange string,queue string) (<-chan amqp.Delivery,error) {
	ch, err := Conn.Channel()
	defer ch.Close()
	if err != nil {
		return nil,err
	}
	err = ch.ExchangeDeclare(exchange,"fanout",true,false,false,false,nil)
	if err != nil {
		panic(err)
	}
	q, err := ch.QueueDeclare(queue,false,false,true,false,nil)
	if err != nil {
		return nil,err
	}
	err = ch.QueueBind(q.Name,"",exchange,false,nil)
	if err != nil {
		return nil,err
	}
	msgs, err := ch.Consume(q.Name,"",false,false,false,false,nil)
	if err != nil {
		return nil,err
	}
	return msgs,nil
}
/**
路由模式  生产者
 */
func ProtoDirect(exchange string,body string,routing string) (bool,error) {
	ch, err := Conn.Channel()
	defer ch.Close()
	if err != nil {
		return false, err
	}
	err = ch.ExchangeDeclare(exchange,"direct",true,false,false,false,nil)
	if err != nil {
		return false, err
	}
	err = ch.Publish(exchange,routing,false,false,amqp.Publishing{
			Body:        []byte(body),
		})
	if err != nil {
		return false, err
	}
	return true,nil
}

