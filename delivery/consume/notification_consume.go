package consume

import (
	csmeService "notif-engine/service"
)


func NewNotificationConsume(consume csmeService.ConsumeNotificationService) {
	FirebaseConsume()
	EmailConsume()
	EmailConsumeArtikel(consume)
}

func FirebaseConsume(){
	
}

func EmailConsume(){
	
}

func EmailConsumeArtikel(){

}