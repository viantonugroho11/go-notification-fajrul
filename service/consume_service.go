package service

import (
	"fmt"
	"notif-engine/model"
	"notif-engine/repository"
)

type ConsumeNotificationService interface {
	ConsumeNotificationFirebase() (result model.PayloadNotificationRequest, err error)
	ConsumeNotificationEmailArtikel(topicName string) (result model.PayloadNotificationRequest, err error)
}

type consumeNotificationService struct {
	msBroker repository.MessageBrokerNotificationRepository
}

func NewConsumeNotificationService(msBroker repository.MessageBrokerNotificationRepository) ConsumeNotificationService {
	return &consumeNotificationService{msBroker}
}

func (s *consumeNotificationService) ConsumeNotificationFirebase() (result model.PayloadNotificationRequest, err error) {	
	result , err = s.msBroker.ConsumeNotificationFirebase()
	if err != nil {
		return result, err
	}
	
  return result, nil	
}

func (s *consumeNotificationService) ConsumeNotificationEmailArtikel(topicName string) (result model.PayloadNotificationRequest, err error) {
	declareQueue := s.msBroker.QueueDeclareRepo(topicName)
	fmt.Println("declareQueue", declareQueue)
	consume , err := s.msBroker.ConsumeNotifArtikel(declareQueue)
	if err != nil {
		return result, err
	}
	fmt.Println("consume", consume)
	go func() {
		s.msBroker.ConsumeWorkerEmail(consume)
	}()
	return result, nil
}