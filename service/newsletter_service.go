package service

import (
	"context"
	"notif-engine/model"
	"notif-engine/repository"
)


type NewsletterService interface {
		GetAllNewsletter(ctx context.Context) (result []model.GetAllNewsletter, err error)
}

type newsletterService struct {
	msBroker repository.MysqlNewsletterRepository
}

func NewNewsletterService(msNewsletter repository.MysqlNewsletterRepository) NewsletterService {
	return &newsletterService{msNewsletter}
}

func (s *newsletterService) GetAllNewsletter(ctx context.Context) (result []model.GetAllNewsletter, err error) {
	
	result, err = s.msBroker.GetAllNewsletter()
	if err != nil {
		return nil, err
	}

	return result, nil
}