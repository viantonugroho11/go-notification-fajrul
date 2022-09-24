package service

import (
	"context"
	"notif-engine/model"
	"notif-engine/repository"
)

type ConsumeNotificationService interface {
	ConsumeNotificationFirebase() (result model.PayloadNotificationRequest, err error)
	ConsumeNotificationEmailArtikel(ctx context.Context, topicName string) (result model.PayloadNotificationRequest, err error)
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

func (s *consumeNotificationService) ConsumeNotificationEmailArtikel(ctx context.Context,topicName string) (result model.PayloadNotificationRequest, err error) {
	declareQueue := s.msBroker.QueueDeclareRepo(topicName)
	consume , err := s.msBroker.ConsumeNotifArtikel(ctx, declareQueue)
	if err != nil {
		return result, err
	}
	go func() {
		s.msBroker.ConsumeWorkerEmail(ctx, consume)
	}()
	return result, nil
}