package handlers_test

import (
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
	token := verifyResponseBody["token"]

	updateUserRequest := handlers.MakeTestRequest(t, "/auth/updateAccount", handlers.UpdateAccountRequest{
		ID: user.ID,
		Location: models.Location{
			Longitude: 1.2468,
			Latitude:  -1.2468,
		},
	}, "POST", &token)

	getUserResponse := handlers.BootstrapServer(updateUserRequest, routeHandlers)
	responseBody := handlers.DecodeResponse(t, getUserResponse)
	assert.Equal(t, "user successfully updated", responseBody["message"])
}

func TestUpdateUser400(t *testing.T) {
	code, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(user)

	verifyUser := handlers.MakeTestRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": code,
	}, "POST", nil)

	verifyResponse := handlers.BootstrapServer(verifyUser, routeHandlers)
	verifyResponseBody := handlers.DecodeResponse(t, verifyResponse)
	token := verifyResponseBody["token"]

	updateUserRequest := handlers.MakeTestRequest(t, "/auth/updateAccount", handlers.UpdateAccountRequest{
		ID: "d84f1416-63f0-422a-95ed-9ce344629ae2",
		Location: models.Location{
			Longitude: 1.2468,
			Latitude:  -1.2468,
		},
	}, "POST", &token)

	updateUserResponse := handlers.BootstrapServer(updateUserRequest, routeHandlers)
	responseBody := handlers.DecodeResponse(t, updateUserResponse)
	assert.Equal(t, "not authorized", responseBody["message"])
}

func TestUpdateUser200RedisCache(t *testing.T) {
	code, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(user)

	verifyUser := handlers.MakeTestRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": code,
	}, "POST", nil)

	verifyResponse := handlers.BootstrapServer(verifyUser, routeHandlers)
	verifyResponseBody := handlers.DecodeResponse(t, verifyResponse)
	token := verifyResponseBody["token"]

	updateUserRequest := handlers.MakeTestRequest(t, "/auth/updateAccount", handlers.UpdateAccountRequest{
		ID: user.ID,
		Location: models.Location{
			Longitude: 1.2468,
			Latitude:  -1.2468,
		},
		Interests: []string{"edging", "camping", "basketball"},
	}, "POST", &token)

	getUserResponse := handlers.BootstrapServer(updateUserRequest, routeHandlers)
	responseBody := handlers.DecodeResponse(t, getUserResponse)
	assert.Equal(t, "user successfully updated", responseBody["message"])
}
