package router

import (
	"github.com/gin-gonic/gin"
	"ai-powered-health-bot/controllers"
)

func Initialize(r *gin.Engine) {
	r.POST("/webhook", controllers.WebhookController)
	r.GET("/webhook", controllers.VerifyWebhook)
}
