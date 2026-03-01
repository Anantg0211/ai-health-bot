package services

import (
	"ai-powered-health-bot/db"
	"ai-powered-health-bot/models"

	"github.com/openai/openai-go"
)

func SaveChat(userID uint, role string, message string) error {

	chat := models.Chat{
		UserID:  userID,
		Role:    role,
		Message: message,
	}

	return db.DB.Create(&chat).Error
}

func GetChatHistory(userID uint) ([]openai.ChatCompletionMessageParamUnion, error) {

	var chats []models.Chat

	err := db.DB.
		Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(10).
		Find(&chats).Error

	if err != nil {
		return nil, err
	}

	var messages []openai.ChatCompletionMessageParamUnion

	for i := len(chats) - 1; i >= 0; i-- {

		c := chats[i]

		if c.Role == "user" {
			messages = append(messages, openai.UserMessage(c.Message))
		} else {
			messages = append(messages, openai.AssistantMessage(c.Message))
		}
	}

	return messages, nil
}