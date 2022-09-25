package service

import (
	"encoding/json"
	"fmt"
	// "fmt"
	"notif-engine/model"
	"notif-engine/repository"
	// "sync"
)

type ConsumeNotificationService interface {

	ConsumeNotificationEmailArtikel(topicName string) (result model.PayloadNotificationRequest, err error)
}

type consumeNotificationService struct {
	msBroker  repository.MessageBrokerNotificationRepository
	emailRepo repository.EmailRepository
}

func NewConsumeNotificationService(msBroker repository.MessageBrokerNotificationRepository, emailRepo repository.EmailRepository) ConsumeNotificationService {
	return &consumeNotificationService{msBroker, emailRepo}
}


func (s *consumeNotificationService) ConsumeNotificationEmailArtikel(topicName string) (result model.PayloadNotificationRequest, err error) {
	declareQueue := s.msBroker.QueueDeclareRepo(topicName)
	consume, _ := s.msBroker.ConsumeNotifArtikel(declareQueue)

	go func(){
		for d := range consume {
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
			_,err=s.emailRepo.EmailPushRepo(&result)
			if err!=nil{
				fmt.Println(err)
			}

	}
	}()
	<-consume
	return result, nil
}
