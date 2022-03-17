package handlers_test

import (
	"testing"
	"time"

	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestVerifyOTPEndpoint_HappyPath(t *testing.T) {
	code, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)

	req := MakeRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": code,
	})

	response := BootstrapServer(req, routeHandlers)
	responseBody := DecodeResponse(t, response)
	assert.Equal(t, "otp code verified", responseBody["message"])
}

func TestVerifyOTPEndpoint_UnhappyPath(t *testing.T) {
	f := faker.New()

	_, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)

	req := MakeRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": f.Numerify("#####"),
	})

	response := BootstrapServer(req, routeHandlers)
	responseBody := DecodeResponse(t, response)
	assert.Equal(t, "failed to verify otp", responseBody["message"])
}

func TestVerifyOTPEndpoint_UnHappyPathNoOTP(t *testing.T) {
	var otp string
	_, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)

	req := MakeRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": otp,
	})

	response := BootstrapServer(req, routeHandlers)
	responseBody := DecodeResponse(t, response)
	assert.Equal(t, "must provide an OTP and an ID. fields cannot be left empty", responseBody["message"])
}

func TestVerifyOTPEndpoint_UnHappyPathExpiredOTP(t *testing.T) {
	code, _, user := handlers.CreateTestUser(t)
	user.OTPCreatedAt = user.OTPCreatedAt.Add(-5 * time.Minute)
	handlers.SeedDB(&user)

	req := MakeRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": code,
	})

	response := BootstrapServer(req, routeHandlers)
	responseBody := DecodeResponse(t, response)
	assert.Equal(t, "otp has expired. regenerate a new one", responseBody["message"])
}
