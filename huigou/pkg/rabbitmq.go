package pkg

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
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
func CusExc(exchange string,queue string)  {
	ch, err := Conn.Channel()
	defer ch.Close()
	if err != nil {
		panic(err)
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
		panic(err)
	}
	rand.Seed(5)
	q, err := ch.QueueDeclare(
		queue,    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		panic(err)
	}

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		exchange, // exchange
		false,
		nil,
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

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
/**
路由模式
 */
func ProtoDirect(exchange string,body string,routing string) (bool,error) {
	ch, err := Conn.Channel()
	defer ch.Close()
	if err != nil {
		return false, err
	}

	err = ch.ExchangeDeclare(
		exchange, // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return false, err
	}
	err = ch.Publish(
		exchange,         // exchange
		routing, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return false, err
	}
	return true,nil
}

func CusDirect(exchange string,queue string,routing string)  {
	ch, err := Conn.Channel()
	defer ch.Close()
	if err != nil {
		panic(111)
	}


	err = ch.ExchangeDeclare(
		exchange, // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		queue,    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,        // queue name
		routing,             // routing key
		exchange, // exchange
		false,
		nil)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}



//错误模式封装
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}