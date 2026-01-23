package services

import (
	"ai-powered-health-bot/helper"
	"context"
	"fmt"

	"github.com/openai/openai-go"
)

const systemPrompt = `
You are a health support assistant.

Rules:
- Do NOT diagnose
- Do NOT prescribe medicine
- Give safe home-care advice
- Be empathetic and supportive
- If symptoms are severe, advise seeing a doctor
- Keep response under 120 words
`

func GetLLMResponse(userText string) (string, error) {
	client := helper.GetOpenAiClient()

	messages := []openai.ChatCompletionMessageParamUnion{}
	messages = append(messages, openai.SystemMessage(systemPrompt))
	messages = append(messages, openai.UserMessage(fmt.Sprintf("User Message: %s", userText)))

	resp, err := client.Chat.Completions.New(
		context.Background(),
		openai.ChatCompletionNewParams{
			Model: openai.ChatModelGPT4oMini,
			Messages: messages,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}