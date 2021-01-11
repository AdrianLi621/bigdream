package pkg

import (
	"github.com/streadway/amqp"
)


var conn *amqp.Connection

/**
实例化mq
*/
func NewMQ() *amqp.Connection {
	once.Do(func() {
		conn, err = amqp.Dial("amqp://admin:admin@dev.com:5672/")
		if err != nil {
			panic(err)
		}
	})
	return conn
}
