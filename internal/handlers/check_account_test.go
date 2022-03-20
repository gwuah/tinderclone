package handlers_test

import (
	"testing"

	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestCheckUser_HappyPath(t *testing.T) {
	code, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)

	req := handlers.MakeTestRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": code,
	})

	response := handlers.BootstrapServer(req, routeHandlers)
	responseBody := handlers.DecodeResponse(t, response)
	assert.Equal(t, "otp code verified", responseBody["message"])
	//TODO create user >> verify user >> pump in data and verify count
}
