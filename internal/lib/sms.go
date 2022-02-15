package lib

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type SMS struct {
	Key      string
	SenderID string
}

type Message struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Sms     string `json:"sms"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	ApiKey  string `json:"api_key"`
}

type Response struct {
	MessageId string  `json:"message_id"`
	Message   string  `json:"message"`
	Balance   float64 `json:"balance"`
	User      string  `json:"user"`
}

const API_ENDPOINT = "https://api.ng.termii.com/api/sms/send"

func MakettpPOSTRequest(endpoint string, requestBody interface{}) (*http.Response, error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func NewSMS(apiKey string) *SMS {
	if apiKey == "" {
		log.Fatal("API key required")
	}
	return &SMS{apiKey, "Tinder Clone"}
}

func (s *SMS) SendTextMessage(message Message) (*http.Response, error) {
	message.ApiKey = s.Key
	message.Type = "plain"
	message.Channel = "generic"
	message.From = s.SenderID

	resp, err := MakettpPOSTRequest(API_ENDPOINT, message)
	return resp, err
}
