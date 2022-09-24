package config

import ( 
"log"
	"github.com/appleboy/go-fcm"
)
type FirebaseConfig struct {
	FCMClientID string 
	Client *fcm.Client
}


func AuthFirebase(conf Config)FirebaseConfig{
	client, err := fcm.NewClient(conf.Firebase.FCMClientID)
	if err != nil {
			log.Fatalln(err)
	}
	return FirebaseConfig{
		Client: client,
	}
}