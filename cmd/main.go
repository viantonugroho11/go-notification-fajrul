package main

import (
	"notif-engine/repository"
	"notif-engine/service"
	conf "notif-engine/config"
	httpHandler "notif-engine/delivery/http"
	"github.com/labstack/echo/v4"
	// amqp "github.com/rabbitmq/amqp091-go"
)
// var ch *amqp.Channel

func main() {
	config := conf.New()

	confQueue := conf.InitQueue(config)
	msRepo := repository.NewMessageBrokerRepository(confQueue)

	msBroker := service.NewPublishService(msRepo)
	e := echo.New()
	api := e.Group("/api")
	e.GET("/health",func(c echo.Context) error {
		return c.JSON(200, "OK")
	})
	

	httpHandler.NewNotificationHandler(api.Group("/v1/notification"), msBroker)
	// api
	e.Logger.Fatal(e.Start(":1324"))

}