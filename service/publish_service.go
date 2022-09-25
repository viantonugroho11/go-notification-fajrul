package service

import (
	"context"
	// "fmt"

	"notif-engine/common"
	"notif-engine/model"
	"notif-engine/repository"
)

type PublishService interface {
	PublishNotif(ctx context.Context, user *model.PayloadNotificationRequest) (result *model.ResponseNotification, err error)
	PublishNotificationArtikel(ctx context.Context, user *model.PayloadNotificationArtikel) (result *model.ResponseNotificationBlast, err error)
	PublishNotificationKabarDonasi(ctx context.Context, user *model.PayloadNotificationKabarDonasi) (result *model.ResponseNotificationBlast, err error)
}

type publishService struct {
	msBroker repository.MessageBrokerNotificationRepository
}

func NewPublishService(msBroker repository.MessageBrokerNotificationRepository) PublishService {
	return &publishService{msBroker}
}

func (s *publishService) PublishNotif(ctx context.Context, user *model.PayloadNotificationRequest) (result *model.ResponseNotification, err error) {

	queueDeclare := s.msBroker.QueueDeclareRepo(common.FirebaseKey)

	_, err = s.msBroker.PublishNotifArtikel(ctx, user, queueDeclare)
	if err != nil {
		return nil, err
	}

	result = &model.ResponseNotification{
		Device : user.Device,
		Title: user.Title,
		Type: user.Type,
	}

	return result, nil
}

func (s *publishService) PublishNotificationArtikel(ctx context.Context, user *model.PayloadNotificationArtikel) (result *model.ResponseNotificationBlast, err error) {

	// queueDeclare := s.msBroker.QueueDeclareRepo(common.EmailKey)

	// _, err = s.msBroker.PublishNotifArtikel(ctx, user, queueDeclare)
	// if err != nil {
	// 	return nil, err
	// }

	result = &model.ResponseNotificationBlast{
		Title: user.Title,
		Body: user.Body,
	}

	return result, nil
}

func (s *publishService) PublishNotificationKabarDonasi(ctx context.Context, user *model.PayloadNotificationKabarDonasi) (result *model.ResponseNotificationBlast, err error) {

	// queueDeclare := s.msBroker.QueueDeclareRepo(common.EmailKey)

	// _, err = s.msBroker.PublishNotifArtikel(ctx, user, queueDeclare)
	// if err != nil {
	// 	return nil, err
	// }

	result = &model.ResponseNotificationBlast{
		Title: user.Title,
		Body: user.Body,
	}

	return result, nil
}