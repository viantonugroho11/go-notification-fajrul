package repository

import "notif-engine/model"



func (em *emailRepository) EmailPushRepo(user *model.PayloadNotificationRequest) (string, error){
	return "success", nil
}

