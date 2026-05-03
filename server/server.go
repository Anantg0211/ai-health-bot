package server

import (
	"ai-powered-health-bot/config"
	"ai-powered-health-bot/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	conf := config.GetConfig()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"*"},
	}))
	router.Initialize(r)
	port := conf.GetString("server.port")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
