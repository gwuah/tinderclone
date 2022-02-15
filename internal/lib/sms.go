package lib

import (
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

	resp, err := MakeSMSRequest(API_ENDPOINT, message)
	return resp, err
}
