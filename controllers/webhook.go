package controllers

import (
	"ai-powered-health-bot/services"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type WhatsAppPayload struct {
	Entry []struct {
		Changes []struct {
			Value struct {
				Messages []struct {
					From string `json:"from"`
					Text struct {
						Body string `json:"body"`
					} `json:"text"`
				} `json:"messages"`
			} `json:"value"`
		} `json:"changes"`
	} `json:"entry"`
}


func WebhookController(c *gin.Context) {
	var payload WhatsAppPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": "invalid payload"})
		return
	}

	fmt.Printf("Webhook payload: %+v\n", payload)

	if len(payload.Entry) == 0 || len(payload.Entry[0].Changes) == 0 || len(payload.Entry[0].Changes[0].Value.Messages) == 0 {
		fmt.Println("Ignoring as no content found.")
		c.JSON(200, gin.H{"status": "ignored"})
		return
	}

	msg := payload.Entry[0].Changes[0].Value.Messages[0]
	phone := msg.From
	text := msg.Text.Body

	user, _ := services.GetOrCreateUser(phone)

	history, _ := services.GetChatHistory(user.ID)

	response, err := services.GetLLMResponse(text, history)
	if err != nil {
		fmt.Printf("Error getting LLM response: %v\n", err)
		response = "Sorry, I'm having trouble processing your request right now."
	}

	services.SaveChat(user.ID, "user", text)
	services.SaveChat(user.ID, "assistant", response)

	er := services.SendMessage(phone, response)
	if er != nil {
		fmt.Printf("Error sending WhatsApp message: %v\n", er)
		c.JSON(500, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(200, gin.H{"status": "message sent"})
}

func VerifyWebhook(c *gin.Context) {
	mode := c.Query("hub.mode")
	token := c.Query("hub.verify_token")
	challenge := c.Query("hub.challenge")

	if mode == "subscribe" && token == os.Getenv("VERIFY_TOKEN") {
		c.String(200, challenge)
		return
	}
	c.AbortWithStatus(403)
}

func Test(c *gin.Context) {
	c.String(200, "AI Powered Health Bot is running.")
}