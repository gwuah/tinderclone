package handlers_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestVerifyOTPEndpoint_HappyPath(t *testing.T) {
	code, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)
	verifyReq := map[string]interface{}{
		"id":  user.ID,
		"otp": code,
	}
	verifyResp, verifyErr := handlers.MakeRequest("verifyOTP", os.Getenv("PORT"), verifyReq)
	assert.NoError(t, verifyErr)
	defer verifyResp.Body.Close()

	assert.Equal(t, http.StatusOK, verifyResp.StatusCode)

	var o map[string]interface{}
	assert.NoError(t, json.NewDecoder(verifyResp.Body).Decode(&o))
	assert.Equal(t, "otp code verified", o["message"])
}

func TestVerifyOTPEndpoint_UnhappyPath(t *testing.T) {
	f := faker.New()

	_, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)
	verifyReq := map[string]interface{}{
		"id":  user.ID,
		"otp": f.Numerify("#####"),
	}
	verifyResp, verifyErr := handlers.MakeRequest("verifyOTP", os.Getenv("PORT"), verifyReq)
	assert.NoError(t, verifyErr)
	defer verifyResp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, verifyResp.StatusCode)

	var o map[string]interface{}
	assert.NoError(t, json.NewDecoder(verifyResp.Body).Decode(&o))
	assert.Equal(t, "failed to verify otp", o["message"])
}

func TestVerifyOTPEndpoint_UnHappyPathNoOTP(t *testing.T) {
	var otp string
	_, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)
	verifyReq := map[string]interface{}{
		"id":  user.ID,
		"otp": otp,
	}
	verifyResp, verifyErr := handlers.MakeRequest("verifyOTP", os.Getenv("PORT"), verifyReq)
	assert.NoError(t, verifyErr)
	defer verifyResp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, verifyResp.StatusCode)

	var o map[string]interface{}
	assert.NoError(t, json.NewDecoder(verifyResp.Body).Decode(&o))
	assert.Equal(t, "failed to verify otp", o["message"])
}
