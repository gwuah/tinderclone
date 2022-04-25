package handlers_test

import (
	"testing"
	"time"

	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestVerifyOTPEndpoint200(t *testing.T) {
	code, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)

	req := handlers.MakeTestRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": code,
	})

	response := handlers.BootstrapServer(req, routeHandlers)
	responseBody := handlers.DecodeResponse(t, response)
	assert.Equal(t, "otp code verified", responseBody["message"])
}

func TestVerifyOTPEndpoint400(t *testing.T) {
	f := faker.New()

	_, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)

	req := handlers.MakeTestRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": f.Numerify("#####"),
	})

	response := handlers.BootstrapServer(req, routeHandlers)
	responseBody := handlers.DecodeResponse(t, response)
	assert.Equal(t, "failed to verify otp", responseBody["message"])
}

func TestVerifyOTPEndpoint400NoOTP(t *testing.T) {
	var otp string
	_, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)

	req := handlers.MakeTestRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": otp,
	})

	response := handlers.BootstrapServer(req, routeHandlers)
	responseBody := handlers.DecodeResponse(t, response)
	assert.Equal(t, "must provide an OTP and an ID. fields cannot be left empty", responseBody["message"])
}

func TestVerifyOTPEndpoint400ExpiredOTP(t *testing.T) {
	code, _, user := handlers.CreateTestUser(t)
	user.OTPCreatedAt = user.OTPCreatedAt.Add(-5 * time.Minute)
	handlers.SeedDB(&user)

	req := handlers.MakeTestRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": code,
	})

	response := handlers.BootstrapServer(req, routeHandlers)
	responseBody := handlers.DecodeResponse(t, response)
	assert.Equal(t, "otp has expired. regenerate a new one", responseBody["message"])
}
