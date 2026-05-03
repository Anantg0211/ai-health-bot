package services

import (
	"ai-powered-health-bot/config"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
	conf := config.GetConfig()
	url := fmt.Sprintf(
		conf.GetString("waba.api_url"),
		conf.GetString("waba.wa_number_id"),
	)

	fmt.Printf("Whatsapp url: %s\n", url)

	reqBody := MessageRequest{
		MessagingProduct: "whatsapp",
		To:               phone,
		Type:             "text",
	}

	reqBody.Text.Body = message

	fmt.Printf("Whatsapp request body: %+v\n", reqBody)

	jsonData, _ := json.Marshal(reqBody)

	fmt.Printf("Whatsapp JSON data: %s\n", string(jsonData))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+conf.GetString("waba.wa_token"))

	fmt.Printf("Whatsapp request: %v\n", req)

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