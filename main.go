package main

import (
	"os"
	"ai-powered-health-bot/config"
	"ai-powered-health-bot/server"
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
	config.Init(environment, "")
	server.Start()
}
