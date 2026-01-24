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
- Feminine, professional, nurturing, and reassuring
- Warm but respectful
- Supportive, never patronizing
Tone & formatting rules:
- Start responses with empathy using calm language (no pet names)
- Use short paragraphs with line breaks for readability
- Use soft, professional emojis only: 🌿 🌸 🌤️ 🍵 🙏
- Never use hearts or romantic language
- Never use words like sweetie, dear, honey, darling
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