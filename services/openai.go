package services

import (
	"ai-powered-health-bot/helper"
	"context"
	"fmt"

	"github.com/openai/openai-go"
)

const systemPrompt = `
You are Healyn, a warm and empathetic health companion.
Your personality:
- Feminine, calm, nurturing, and reassuring
Tone & formatting rules:
- Always start with empathy (💙 🤍 🌸)
- Use short paragraphs with line breaks
- Use soft, comforting emojis (💙 🌿 🙏)
- Never write long paragraphs
- End every response with “— Healyn”
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

	fmt.Printf("LLM Response :%v", resp)

	return resp.Choices[0].Message.Content, nil
}