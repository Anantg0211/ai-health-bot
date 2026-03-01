package services

import (
	"fmt"

	"github.com/openai/openai-go"
)

type Tool func(string, []openai.ChatCompletionMessageParamUnion) (string, error)

var tools = map[string]Tool{
	"triage_emergency": EmergencyTriageTool,
	"mental_support":   MentalSupportTool,
	"symptom_guidance": SymptomGuidanceTool,
	"first_aid":        FirstAidTool,
	"chat":             ChatTool,
}

func EmergencyTriageTool(userText string, history []openai.ChatCompletionMessageParamUnion) (string, error) {

	context := fmt.Sprintf(`
User message: %s

Respond as Healyn.

If the symptom may be serious (like chest pain), advise seeking immediate medical help while staying calm.
`, userText)

	resp, err := BuildRequest(context, healynPrompt, history)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func MentalSupportTool(userText string, history []openai.ChatCompletionMessageParamUnion) (string, error) {

	context := fmt.Sprintf(`
User message: %s

Provide calming emotional support and simple grounding advice.
`, userText)

	resp, err := BuildRequest(context, healynPrompt, history)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func SymptomGuidanceTool(userText string, history []openai.ChatCompletionMessageParamUnion) (string, error) {

	context := fmt.Sprintf(`
User message: %s

Provide safe home-care guidance and suggest seeing a doctor if symptoms worsen.
`, userText)

	resp, err := BuildRequest(context, healynPrompt, history)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func FirstAidTool(userText string, history []openai.ChatCompletionMessageParamUnion) (string, error) {

	context := fmt.Sprintf(`
User message: %s

Provide basic first aid guidance for minor injuries.
`, userText)

	resp, err := BuildRequest(context, healynPrompt, history)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func ChatTool(userText string, history []openai.ChatCompletionMessageParamUnion) (string, error) {

	resp, err := BuildRequest(userText, healynPrompt, history)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}