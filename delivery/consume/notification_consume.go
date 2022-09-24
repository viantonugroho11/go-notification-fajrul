package consume

import (
	// "fmt"
	"fmt"
	"notif-engine/common"
	csmeService "notif-engine/service"
)

func NewNotificationConsume(consume csmeService.ConsumeNotificationService) {
	go EmailConsumeArtikel(consume)
	fmt.Println("masuk")
}

func EmailConsumeArtikel(consume csmeService.ConsumeNotificationService) {
		result, err := consume.ConsumeNotificationEmailArtikel(common.FirebaseKey)
		if err != nil {
			return json.NewEncoder(w).Encode(err)
		}
		fmt.Println("result", result)
}
