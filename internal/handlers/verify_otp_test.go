package handlers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestVerifyOTPEndpoint_HappyPath(t *testing.T) {
	code, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)
	requestPostBody := map[string]interface{}{
		"id":  user.ID,
		"otp": code,
	}

	body, err := json.Marshal(requestPostBody)
	if err != nil {
		log.Print(err)
	}
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/verifyOTP", bytes.NewReader(body))
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()

	testServer.ServeHTTP(responseRecorder, req)

	var responseBody map[string]interface{}
	assert.NoError(t, json.NewDecoder(responseRecorder.Result().Body).Decode(&responseBody))
	assert.Equal(t, "otp code verified", responseBody["message"])
}

func TestVerifyOTPEndpoint_UnhappyPath(t *testing.T) {
	f := faker.New()

	_, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)
	requestPostBody := map[string]interface{}{
		"id":  user.ID,
		"otp": f.Numerify("#####"),
	}
	body, err := json.Marshal(requestPostBody)
	if err != nil {
		log.Print(err)
	}
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/verifyOTP", bytes.NewReader(body))
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()

	testServer.ServeHTTP(responseRecorder, req)

	var responseBody map[string]interface{}
	assert.NoError(t, json.NewDecoder(responseRecorder.Result().Body).Decode(&responseBody))
	assert.Equal(t, "failed to verify otp", responseBody["message"])
}

func TestVerifyOTPEndpoint_UnHappyPathNoOTP(t *testing.T) {
	var otp string
	_, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)

	requestPostBody := map[string]interface{}{
		"id":  user.ID,
		"otp": otp,
	}

	body, err := json.Marshal(requestPostBody)
	if err != nil {
		log.Print(err)
	}
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/verifyOTP", bytes.NewReader(body))
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()

	testServer.ServeHTTP(responseRecorder, req)

	var responseBody map[string]interface{}
	assert.NoError(t, json.NewDecoder(responseRecorder.Result().Body).Decode(&responseBody))
	assert.Equal(t, "must provide an OTP and an ID. fields cannot be left empty", responseBody["message"])
}

func TestVerifyOTPEndpoint_UnHappyPathExpiredOTP(t *testing.T) {
	code, _, user := handlers.CreateTestUser(t)
	user.OTPCreatedAt = user.OTPCreatedAt.Add(-5 * time.Minute)
	handlers.SeedDB(&user)
	requestPostBody := map[string]interface{}{
		"id":  user.ID,
		"otp": code,
	}

	body, err := json.Marshal(requestPostBody)
	if err != nil {
		log.Print(err)
	}
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/verifyOTP", bytes.NewReader(body))
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()

	testServer.ServeHTTP(responseRecorder, req)

	var responseBody map[string]interface{}
	assert.NoError(t, json.NewDecoder(responseRecorder.Result().Body).Decode(&responseBody))
	assert.Equal(t, "otp has expired. regenerate a new one", responseBody["message"])
}
