package lib

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Termii struct {
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

func NewTermii(apiKey string) (*Termii, error) {
	if apiKey == "" {
		return nil, errors.New("API key required")
	}
	return &Termii{APIKey: apiKey, SenderID: "electra"}, nil
}

func (t *Termii) SendTextMessage(to string, sms string) (Response, error) {
	m := Message{
		ApiKey:  t.APIKey,
		From:    t.SenderID,
		Type:    "plain",
		Channel: "generic",
		To:      to,
		Sms:     sms,
	}

	response := Response{}

	payload, err := json.Marshal(&m)
	if err != nil {
		return Response{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		"https://api.ng.termii.com/api/sms/send",
		strings.NewReader(string(payload)),
	)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return Response{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, err
	}

	return response, err
}
