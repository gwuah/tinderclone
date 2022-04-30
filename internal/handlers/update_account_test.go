package handlers_test

import (
	//"fmt"
	"testing"

	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/gwuah/tinderclone/internal/models"
	"github.com/stretchr/testify/assert"
)



func TestUpdateUser200(t *testing.T) {
	code, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(user)

	verifyUser := handlers.MakeTestRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": code,
	}, "POST", nil)

	verifyResponse := handlers.BootstrapServer(verifyUser, routeHandlers)
	verifyResponseBody := handlers.DecodeResponse(t, verifyResponse)
	_ = verifyResponseBody["token"]

	getUserRequest := handlers.MakeTestRequest(t, "/updateAccount" , handlers.UpdateAccountRequest{
		ID: user.ID,
		Location: models.Location{
			Longitude: 1.232,
			Latitude: -1.232,
		},
	}, "POST", nil)

	getUserResponse := handlers.BootstrapServer(getUserRequest, routeHandlers)
	responseBody := handlers.DecodeResponse(t, getUserResponse)
	assert.Equal(t, "user successfully retrieved", responseBody["message"])
}