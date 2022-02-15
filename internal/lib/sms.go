package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type SMS struct {
	APIKey   string
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

func NewSMS(senderID, apiKey string) (*SMS, error) {
	if apiKey == "" {
		return nil, errors.New("API key required")
	}
	return &SMS{APIKey: apiKey, SenderID: senderID}, nil
}

func (s *SMS) SendTextMessage(to string, sms string) (Response, error) {
	message := Message{
		ApiKey:  s.APIKey,
		From:    s.SenderID,
		Type:    "plain",
		Channel: "generic",

		To:  to,
		Sms: sms,
	}

	_, err := MakeHttpPOSTRequest(API_ENDPOINT, message)

	return Response{}, err
}

func MakeHttpPOSTRequest(endpoint string, requestBody interface{}) (*http.Response, error) {
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
