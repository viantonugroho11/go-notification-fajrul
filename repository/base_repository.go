package repository

import (
	"context"
	// "database/sql"
	"notif-engine/config"
	"notif-engine/model"

	amqp "github.com/rabbitmq/amqp091-go"
)

type postgreRepository struct {
	Ch *amqp.Channel
}

func NewPostgreRepository(ch *amqp.Channel) PostgreUserRepository {
	return &postgreRepository{ch}
}

type PostgreUserRepository interface {
	
}

type messageRepository struct {
	confQueue config.MessageBrokerConfig
}

func NewMessageBrokerRepository(confQueue config.MessageBrokerConfig) MessageBrokerNotificationRepository {
	return &messageRepository{confQueue}
}

type MessageBrokerNotificationRepository interface {
	// PublishNotification(ctx context.Context, user *model.PayloadNotificationRequest) (string, error)

	
	PublishNotifArtikel(ctx context.Context, user *model.PayloadNotificationRequest, channel *amqp.Queue) (string, error)
	
	
	QueueDeclareRepo(name string) *amqp.Queue

	//deprecated
	// ConsumeNotificationFirebase() (result model.PayloadNotificationRequest, err error)

	ConsumeNotifArtikel(channel *amqp.Queue)(result <-chan amqp.Delivery, err error)
	ConsumeWorkerEmail(message <-chan amqp.Delivery)
	
}

type firebaseRepository struct {
	confQueue config.MessageBrokerConfig
	confFb 	config.FirebaseConfig
}

func NewFirebaseRepository(confQueue config.MessageBrokerConfig, confFb config.FirebaseConfig) FirebaseRepository {
	return &firebaseRepository{confQueue, confFb}
}

type FirebaseRepository interface {
	FirebasePushRepo(user *model.PayloadNotificationRequest) (string, error)
}

type emailRepository struct {
	confQueue config.MessageBrokerConfig
	confEmail config.EmailConfig
}

func NewEmailRepository(confQueue config.MessageBrokerConfig, confEmail config.EmailConfig) EmailRepository {
	return &emailRepository{confQueue, confEmail}
}

type EmailRepository interface {
	EmailPushRepo(user *model.PayloadNotificationRequest) (string, error)
}