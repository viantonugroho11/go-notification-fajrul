package consume

import (
	"fmt"
	"notif-engine/common"
	csmeService "notif-engine/service"
)

func NewNotificationConsume(consume csmeService.ConsumeNotificationService) {
	go EmailConsumeArtikel(consume)
	go EmailConsumeNewsletterArtikel(consume)
	go fmt.Println("masuk ke consume")
}

func EmailConsumeArtikel(consume csmeService.ConsumeNotificationService) {
		_, err := consume.ConsumeNotificationEmailArtikel(common.FirebaseKey)
		if err != nil {
			fmt.Println("error consume", err)
		}
}

func EmailConsumeNewsletterArtikel(consume csmeService.ConsumeNotificationService) {
	_, err := consume.ConsumeEmailNewsletterArtikelService(common.EmailArtikel)
	fmt.Print("Email Artikel")
	if err != nil {
		fmt.Println("error consume", err)
	}
}

func EmailConsumeKabarDonasi(consume csmeService.ConsumeNotificationService){
	_, err := consume.ConsumeEmailKabarDonasiService(common.EmailKabarDonasi)
	if err != nil {
		fmt.Println("error consume", err)
	}
}