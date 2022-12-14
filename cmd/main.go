package main

import (
	"fmt"
	conf "notif-engine/config"
	consumeHandler "notif-engine/delivery/consume"
	httpHandler "notif-engine/delivery/http"
	"notif-engine/repository"
	"notif-engine/service"

	"os"
	"github.com/labstack/echo/v4"
	// amqp "github.com/rabbitmq/amqp091-go"
)

// var ch *amqp.Channel

func main() {
	config := conf.New()
	conf.InitHeroku()

	confQueue := conf.InitQueue(config)
	confEmail := conf.AuthEmail(config)
	confDb := conf.InitDb(config)

	msRepo := repository.NewMessageBrokerRepository(confQueue)
	emailRepo := repository.NewEmailRepository(confQueue, confEmail)
	mysqlNewsRepo := repository.NewMysqlNewsletterRepository(confDb)
	mysqlKabarDonasiRepo := repository.NewMysqlKabarDonasiRepository(confDb)

	msBroker := service.NewPublishService(msRepo)
	msConsume := service.NewConsumeNotificationService(msRepo, emailRepo,mysqlNewsRepo, mysqlKabarDonasiRepo)
	mysqlNews := service.NewNewsletterService(mysqlNewsRepo)
	
	e := echo.New()
	api := e.Group("/api")
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	fmt.Println("masuk")

	go consumeHandler.NewNotificationConsume(msConsume)
	
	httpHandler.NewNotificationHandler(api.Group("/v1/notification"), msBroker)
	httpHandler.NewNewsletterHandler(api.Group("/v1/newsletter"), mysqlNews)
	// api

	fmt.Println("port: ",os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(":"+os.Getenv("PORT")))

}
