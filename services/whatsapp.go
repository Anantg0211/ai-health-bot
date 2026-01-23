package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type MessageRequest struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             struct {
		Body string `json:"body"`
	} `json:"text"`
}


func SendMessage(phone string, message string) error {
	url := fmt.Sprintf(
		"https://graph.facebook.com/v22.0/%s/messages",
		os.Getenv("WHATSAPP_PHONE_ID"),
	)

	reqBody := MessageRequest{
		MessagingProduct: "whatsapp",
		To:               phone,
		Type:             "text",
	}

	reqBody.Text.Body = message

	jsonData, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("WHATSAPP_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("whatsapp api failed: %d", resp.StatusCode)
	}

	return nil
}