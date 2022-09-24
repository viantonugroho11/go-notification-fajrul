package common

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)


var ch *amqp.Channel

func Manager() *amqp.Channel {
	return ch
}

func DeclareQueue(name string) *amqp.Queue {
	fmt.Println("declare queue1")
	q, _ := ch.QueueDeclare(name, false, false, false, false, nil)
	fmt.Println("declare queue",q)
	return &q
}