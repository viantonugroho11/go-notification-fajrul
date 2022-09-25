package service

import (
	"fmt"
	"notif-engine/model"
	"notif-engine/repository"
	"sync"
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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
			result, err := s.msBroker.ConsumeWorkerEmail(consume)
			if err != nil {
				fmt.Println("error consume", err)
			
			}
			s.emailRepo.EmailPushRepo(result)
		wg.Done()
		}()
	wg.Wait()
	return result, nil
}
