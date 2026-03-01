package services

import (
	"ai-powered-health-bot/helper"
	"context"
	"encoding/json"
	"strings"

	"github.com/openai/openai-go"
)

const healynPrompt = `
You are Healyn, a warm and empathetic health companion.
Your personality:
- Feminine, professional, nurturing, and reassuring
- Warm but respectful
- Supportive, never patronizing
Tone & formatting rules:
- Start responses with empathy using calm language (no pet names)
- Use short paragraphs with line breaks for readability
- Use soft, professional emojis which are in given array [🌿, 🌸, 🌤️, 🍵]
- Never use hearts or romantic language
- Never use words like sweetie, dear, honey, darling
- Never write long paragraphs
- End every response with “— Healyn”
`

func GetLLMResponse(userText string) (string, error) {
	decision, err := DecideAction(userText)
	if err != nil {
		return "", err
	}	

	tool, exists := tools[decision.Action]
	if !exists {
		tool = ChatTool
	}

	return tool(userText)
}

const decisionPrompt = `
You are Healyn's decision engine.

Based on the user's message, choose the correct action.

Actions:
- triage_emergency → serious symptoms like chest pain, breathing issues
- mental_support → anxiety, stress, emotional distress
- symptom_guidance → common symptoms like headache, mild fever
- first_aid → small injuries like cuts or burns
- chat → normal conversation

Return ONLY valid JSON.

Example:
{"action":"symptom_guidance"}
`

type AgentDecision struct {
	Action string `json:"action"`
}

func DecideAction(userText string) (AgentDecision, error) {
	resp, err := BuildRequest(userText, decisionPrompt)

	if err != nil {
		return AgentDecision{}, err
	}

	content := resp.Choices[0].Message.Content
	content = strings.TrimSpace(content)
	
	var decision AgentDecision
	err = json.Unmarshal([]byte(content), &decision)

	return decision, err
}

func BuildRequest(userText string, prompt string) (*openai.ChatCompletion, error) {
	client := helper.GetOpenAiClient()

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(prompt),
		openai.UserMessage(userText),
	}

	resp, err := client.Chat.Completions.New(
		context.Background(),
		openai.ChatCompletionNewParams{
			Model: openai.ChatModelGPT4oMini,
			Messages: messages,
		},
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}