package config

import "net/smtp"

type EmailConfig struct {
	Host      string
	Port      int
	Username  string
	Password  string
	SmtpAuth smtp.Auth
}

func AuthEmail(conf Config) EmailConfig{

	auth := smtp.PlainAuth("", conf.Email.Email, conf.Email.Password, conf.Email.Host)
	return EmailConfig{
		Host:     conf.Email.Host,
		Port:     conf.Email.Port,
		Username: conf.Email.Email,
		Password: conf.Email.Password,
		SmtpAuth: auth,
	}
}
