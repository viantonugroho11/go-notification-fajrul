package service

import (
	"context"
	// "fmt"

	"notif-engine/common"
	"notif-engine/model"
	"notif-engine/repository"
)

type PublishService interface {
	PublishNotif(ctx context.Context, user *model.PayloadNotificationRequest) (string, error)
}

type publishService struct {
	msBroker repository.MessageBrokerNotificationRepository
}

func NewPublishService(msBroker repository.MessageBrokerNotificationRepository) PublishService {
	return &publishService{msBroker}
}

func (s *publishService) PublishNotif(ctx context.Context, user *model.PayloadNotificationRequest) (string, error) {

	queueDeclare := s.msBroker.QueueDeclareRepo(common.FirebaseKey)

	data, err := s.msBroker.PublishNotifArtikel(ctx, user, queueDeclare)
	if err != nil {
		return "", err
	}
	
	return data, nil
}
// func (s *publishService) PublishNotif(ctx context.Context, user &model.PayloadNotificationRequest) (string, error) {
// 	return nil, nil
// }