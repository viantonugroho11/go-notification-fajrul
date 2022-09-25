package repository

import (
	"fmt"
	"net/smtp"
	"notif-engine/model"
)

func (em *emailRepository) EmailPushRepo(user *model.PayloadNotificationRequest) (string, error) {
	to := []string{
		user.Device,
	}
	msg := []byte("From: " + em.confEmail.FromEmail + "\r\n" +
		"To: " + user.Device + "\r\n" +
		"Subject: " + user.Title + "\r\n\r\n" +
		"" + user.Body + "\r\n")

	err := smtp.SendMail(em.confEmail.AddressMail, em.confEmail.SmtpAuth, em.confEmail.Username, to, msg)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Print("Email Sent!")
	return "success", nil
}

func (em *emailRepository) EmailPushTiketDonasi() {

	//use template html

}

func (em *emailRepository) EmailPushNewsletter() {

	//use template html

}

func (em *emailRepository) EmailPushKabarDonasi() {

	//use template html

}
