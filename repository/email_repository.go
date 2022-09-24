package repository

import (
	"net/smtp"
	"notif-engine/model"
)



func (em *emailRepository) EmailPushRepo(user *model.PayloadNotificationRequest) (string, error){
	to := []string{
		user.Device,
	}

	msg := []byte("From: " + em.confEmail.AddressMail + "\r\n" +
		"To: " + user.Device + "\r\n" +
		"Subject: " + user.Title + "\r\n\r\n" +
		"" + user.Body + "\r\n")
	err := smtp.SendMail(em.confEmail.AddressMail,em.confEmail.SmtpAuth, em.confEmail.AddressMail, to, msg)
	if err != nil {
		return "", err
	}
	return "success", nil
}

func (em *emailRepository) EmailPushTiketDonasi()(){


	//use template html
	
}

func (em *emailRepository) EmailPushNewsletter()(){


	//use template html
	
}

func (em *emailRepository) EmailPushKabarDonasi()(){


	//use template html
	
}