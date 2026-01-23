package server

import (
	"ai-powered-health-bot/config"
	"ai-powered-health-bot/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	config := config.GetConfig()
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.GetStringSlice("cors_allow_origin")
	// corsConfig.AllowHeaders = []string{"Authorization", "content-type"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowMethods = []string{"DELETE", "PUT", "GET", "POST"}
	r.Use(cors.New(corsConfig))
	router.Initialize(r)
	r.Run(config.GetString("server.host") + ":" + config.GetString("server.port"))
}
