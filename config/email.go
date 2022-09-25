package config

import (
	"net/smtp"
)

type EmailConfig struct {
	Host      string
	Port      int
	Username  string
	Password  string
	AddressMail string
	FromEmail string
	SmtpAuth smtp.Auth

}

func AuthEmail(conf Config) EmailConfig{
	auth := smtp.PlainAuth("", conf.Email.Email, conf.Email.Password, conf.Email.Host)
	return EmailConfig{
		Host:     conf.Email.Host,
		Port:     conf.Email.Port,
		Username: conf.Email.Email,
		Password: conf.Email.Password,
		AddressMail: conf.Email.AddressMail,
		FromEmail: conf.Email.FromNameEmail,
		SmtpAuth: auth,
	}
}
