package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestCreateAccountEndpoint(t *testing.T) {
	port := os.Getenv("port")

	PostBody := map[string]interface{}{
		"phone_number": "5rxxxxxxxcscx",
	}

	body, err := json.Marshal(PostBody)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:%s/createAccount", port), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.StatusCode != 201 {
		t.Fatalf("expected %d, got %v", 201, resp.StatusCode)
	}

	var m map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	resp.Body.Close()

	if m["message"] != "user succesfully created." {
		t.Fatalf("expected %s, got %v", "user succesfully created.", m["message"])
	}

}
