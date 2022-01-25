package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func MakeRequest(endpoint string, port string, requestBody interface{}) (*http.Response, error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:%s/%s", port, endpoint), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
