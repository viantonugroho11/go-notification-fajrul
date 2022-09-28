package config

import (
	"github.com/joho/godotenv"
  "log"
)


func InitHeroku(){
err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}