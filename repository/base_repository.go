package repository

import (
	"context"
	// "database/sql"
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

//mysql newsletter
type mysqlNewsletterRepository struct {
	Conn config.DatabaseConfig
}

func NewMysqlNewsletterRepository(conn config.DatabaseConfig) MysqlNewsletterRepository {
	return &mysqlNewsletterRepository{conn}
}

type MysqlNewsletterRepository interface {
	GetAllNewsletter() ([]model.GetAllNewsletter, error)
}

type mysqlKabarDonasiRepository struct {
	Conn config.DatabaseConfig
}

func NewMysqlKabarDonasiRepository(conn config.DatabaseConfig) MysqlKabarDonasiRepository {
	return &mysqlKabarDonasiRepository{conn}
}

type MysqlKabarDonasiRepository interface {
	GetAllUserByDonasiID(ctx context.Context, id string) (result []model.GetEmailUserKabarDonasi,err error)
	GetUserStatusNotyetByDonasiID(ctx context.Context, id string) (result []model.GetEmailUserKabarDonasi,err error)
}


type messageRepository struct {
	confQueue config.MessageBrokerConfig
}

func NewMessageBrokerRepository(confQueue config.MessageBrokerConfig) MessageBrokerNotificationRepository {
	return &messageRepository{confQueue}
}

type MessageBrokerNotificationRepository interface {
	PublishNotifArtikel(ctx context.Context, user *model.PayloadNotificationRequest, channel *amqp.Queue) (string, error)
	
	QueueDeclareRepo(name string) *amqp.Queue

	ConsumeNotifArtikel(channel *amqp.Queue)(result <-chan amqp.Delivery, err error)
	ConsumeWorkerEmail(message <-chan amqp.Delivery)(data *model.PayloadNotificationRequest, err error)
	

	PublishNotificationArtikelRepo(ctx context.Context, user *model.PayloadNotificationArtikel, channel *amqp.Queue) (string, error)

	PublishNotificationKabarDonasiRepo(ctx context.Context, user *model.PayloadNotificationKabarDonasi, channel *amqp.Queue) (string, error)
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