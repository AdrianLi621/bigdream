package pkg

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
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




//错误模式封装
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
