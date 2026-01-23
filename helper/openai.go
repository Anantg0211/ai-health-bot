package helper

import (
	"ai-powered-health-bot/config"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func GetOpenAiClient() *openai.Client {
	config := config.GetConfig()
	openaiClient := openai.NewClient(
		option.WithAPIKey(config.GetString("openai_keys.transcription")),
	)
	return &openaiClient
}