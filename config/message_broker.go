package config

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var ch *amqp.Channel

type MessageBrokerConfig struct {
	Host                 string
	Port                 string
	Username             string
	Password             string
	QueueRoutingFirebase string
	Dial                 string
	Ch                   *amqp.Channel
}

func InitQueue(conf Config) MessageBrokerConfig {
	// Connect to RabbitMQ server
	connectString := fmt.Sprintf(conf.Queue.Dial)
	conn, err := amqp.Dial(connectString) //Insert the  connection string
	if err != nil {
		panic(err)
		// fmt.Errorf("Failed to connect to RabbitMQ: %s", err)
	}
	fmt.Println("Connected to RabbitMQ")
	// helper.FailOnError(err, "RabbitMQ connection failure", "RabbitMQ Connection Established")
	// defer conn.Close()

	//Connect to the channel
	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	// fmt.Println("Connected to Channel")
	return MessageBrokerConfig{
		Host:     conf.Queue.Host,
		Username: conf.Queue.Username,
		Password: conf.Queue.Password,
		Dial:     conf.Queue.Dial,
		Ch:       ch,
	}
}
