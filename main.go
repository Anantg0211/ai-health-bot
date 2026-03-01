package main

import (
	"ai-powered-health-bot/config"
	"ai-powered-health-bot/db"
	"ai-powered-health-bot/models"
	"ai-powered-health-bot/server"
	"os"
)

// @title TruVoice API
// @version 1.0
// @description This is the API documentation for TruVoice
// @contact.name Rizwan Khan
// @contact.email rizwan@trudoc.ae

// @BasePath /

func main() {
	environment := os.Getenv("API_ENV")
	if environment == "" {
		environment = "development"
	}
	
	db.Connect()

	db.DB.AutoMigrate(
		&models.User{},
		&models.Chat{},
	)

	config.Init(environment, "")
	server.Start()
}
