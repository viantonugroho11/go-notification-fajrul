package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"notif-engine/common"
	"notif-engine/model"

	amqp "github.com/rabbitmq/amqp091-go"
)

// var ch *amqp.Channel
//deprecated
func (msBroker *messageRepository) PublishNotification(ctx context.Context, user *model.PayloadNotificationRequest) (string, error) {

	queue, _ := msBroker.confQueue.Ch.QueueDeclare(user.Type, false, false, false, false, nil)

	err := msBroker.confQueue.Ch.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(`{"device":"` + user.Device + `","title":"` + user.Title + `","message":"` + user.Body + `","userid":"` + user.UserID + `"}`),
	})

	if err != nil {
		return "", err
	}
	return common.SuccesMessage, nil
}

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


// deprecated
func (msBroker *messageRepository) ConsumeNotificationFirebase() (result model.PayloadNotificationRequest, err error) {
	queue, _ := msBroker.confQueue.Ch.QueueDeclare(common.FirebaseKey, false, false, false, false, nil)
	_, err = msBroker.confQueue.Ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return result, err
	}
	return result, nil
	// forever := make(chan bool)
}


func (msBroker *messageRepository) ConsumeWorkerEmail(message <-chan amqp.Delivery) {
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
		fmt.Println(result)
		// msBroker.FirebasePushRepo(&result)
	}
	// msBroker.PublishNotification(&result)
}



