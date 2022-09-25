package repository

import (
	"context"
	"encoding/json"
	"notif-engine/common"
	"notif-engine/model"

	amqp "github.com/rabbitmq/amqp091-go"
)


func (msBroker *messageRepository) PublishNotifArtikel(ctx context.Context, user *model.PayloadNotificationRequest, channel *amqp.Queue) (string, error) {
	err := msBroker.confQueue.Ch.Publish("", channel.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(`{"device":"` + user.Device + `","title":"` + user.Title + `","message":"` + user.Body + `","userid":"` + user.UserID + `"}`),
	})

	if err != nil {
		return "", err
	}
	return common.SuccesMessage, nil
}

func (msBroker *messageRepository) ConsumeNotifArtikel(channel *amqp.Queue)(result <-chan amqp.Delivery, err error){
	msgs, err := msBroker.confQueue.Ch.Consume(channel.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}


func (msBroker *messageRepository) QueueDeclareRepo(name string) *amqp.Queue {
	queue, _ := msBroker.confQueue.Ch.QueueDeclare(name, false, false, false, false, nil)
	return &queue
}


func (msBroker *messageRepository) ConsumeWorkerEmail(message <-chan amqp.Delivery)(data *model.PayloadNotificationRequest, err error) {
	var result model.PayloadNotificationRequest
	for d := range message {
		body := string(d.Body)
		jsondata := []byte(body)
		var fire interface{}
		json.Unmarshal(jsondata, &fire)
		decode := fire.(map[string]interface{})
		device := decode["device"].(string)
		userid := decode["userid"].(string)
		message := decode["message"].(string)
		title := decode["title"].(string)
		result = model.PayloadNotificationRequest{
			Device: device,
			UserID: userid,
			Body:   message,
			Title:  title,
		}
		return &result, nil
	}
	return nil, nil
}



