package repository

import (
	"notif-engine/common"
	"notif-engine/model"

	"github.com/appleboy/go-fcm"
)

func (fb *firebaseRepository) FirebasePushRepo(user *model.PayloadNotificationRequest) (string, error) {
	msg := &fcm.Message{
		To: user.Device,
		Notification: &fcm.Notification{
			Title: user.Title,
			Body:  user.Body,
		},
		Data: map[string]interface{}{
			"click_action": "FLUTTER_NOTIFICATION_CLICK",
			"status":       "done",
			"screen":       "home",
		},
	}
	result, err := fb.confFb.Client.Send(msg)
	if result.Failure != 0 {
		return "", err
	}

	// return nil, nil
	return common.SuccesMessage, nil
}