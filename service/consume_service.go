package service

import (
	"fmt"
	"notif-engine/model"
	"notif-engine/repository"
)

type ConsumeNotificationService interface {
	// ConsumeNotificationFirebase() (result model.PayloadNotificationRequest, err error)
	ConsumeNotificationEmailArtikel(topicName string) (result model.PayloadNotificationRequest, err error)
}

type consumeNotificationService struct {
	msBroker repository.MessageBrokerNotificationRepository
	emailRepo repository.EmailRepository
}

func NewConsumeNotificationService(msBroker repository.MessageBrokerNotificationRepository, emailRepo repository.EmailRepository) ConsumeNotificationService {
	return &consumeNotificationService{msBroker,emailRepo}
}

// func (s *consumeNotificationService) ConsumeNotificationFirebase() (result model.PayloadNotificationRequest, err error) {	
// 	result , err = s.msBroker.ConsumeNotificationFirebase()
// 	if err != nil {
// 		return result, err
// 	}
	
//   return result, nil	
// }

func (s *consumeNotificationService) ConsumeNotificationEmailArtikel(topicName string) (result model.PayloadNotificationRequest, err error) {
	declareQueue := s.msBroker.QueueDeclareRepo(topicName)
	fmt.Println("declareQueue", declareQueue)
	consume , err := s.msBroker.ConsumeNotifArtikel(declareQueue)
	if err != nil {
		return result, err
	}
	go func() {
		result,err := s.msBroker.ConsumeWorkerEmail(consume)
		if err != nil {
			fmt.Println("error consume", err)
		}
		s.emailRepo.EmailPushRepo(result)
	}()
	return result, nil
}