package consume

import (
	"fmt"
	"notif-engine/common"
	csmeService "notif-engine/service"
)

func NewNotificationConsume(consume csmeService.ConsumeNotificationService) {
	go EmailConsumeArtikel(consume)
	fmt.Println("masuk")
}

func EmailConsumeArtikel(consume csmeService.ConsumeNotificationService) {
		_, err := consume.ConsumeNotificationEmailArtikel(common.FirebaseKey)
		if err != nil {
			fmt.Println("error consume", err)
		}
}
