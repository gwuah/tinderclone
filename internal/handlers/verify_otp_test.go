package handlers_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/gwuah/tinderclone/internal/handlers"
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
