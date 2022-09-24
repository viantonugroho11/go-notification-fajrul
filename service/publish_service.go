package service

import (
	"context"
	"fmt"

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
	// return s.msBroker.PublishNotif(ctx, user)
	data , err := s.msBroker.PublishNotification(ctx, user)
	if err != nil {
		return "", err
	}
	fmt.Println(data)

	
	return data, nil
}
// func (s *publishService) PublishNotif(ctx context.Context, user &model.PayloadNotificationRequest) (string, error) {
// 	return nil, nil
// }