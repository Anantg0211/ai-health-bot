package services

import "fmt"

type Tool func(string) (string, error)

var tools = map[string]Tool{
	"triage_emergency": EmergencyTriageTool,
	"mental_support":   MentalSupportTool,
	"symptom_guidance": SymptomGuidanceTool,
	"first_aid":        FirstAidTool,
	"chat":             ChatTool,
}

func EmergencyTriageTool(userText string) (string, error) {

	context := fmt.Sprintf(`
User message: %s

Respond as Healyn.

If the symptom may be serious (like chest pain), advise seeking immediate medical help while staying calm.
`, userText)

	resp, err := BuildRequest(context, healynPrompt)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func MentalSupportTool(userText string) (string, error) {

	context := fmt.Sprintf(`
User message: %s

Provide calming emotional support and simple grounding advice.
`, userText)

	resp, err := BuildRequest(context, healynPrompt)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func SymptomGuidanceTool(userText string) (string, error) {

	context := fmt.Sprintf(`
User message: %s

Provide safe home-care guidance and suggest seeing a doctor if symptoms worsen.
`, userText)

	resp, err := BuildRequest(context, healynPrompt)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func FirstAidTool(userText string) (string, error) {

	context := fmt.Sprintf(`
User message: %s

Provide basic first aid guidance for minor injuries.
`, userText)

	resp, err := BuildRequest(context, healynPrompt)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func ChatTool(userText string) (string, error) {

	resp, err := BuildRequest(userText, healynPrompt)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}