package handlers_test

import (
	"fmt"
	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/gwuah/tinderclone/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUser200(t *testing.T) {
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

	updateUserResponse := handlers.BootstrapServer(updateUserRequest, routeHandlers)
	responseBody := handlers.DecodeResponse(t, updateUserResponse)
	assert.Equal(t, "user successfully updated", responseBody["message"])

	getUserRequest := handlers.MakeTestRequest(t, fmt.Sprintf("/auth/getUser/%s", user.ID), map[string]interface{}{}, "GET", &token)

	getUserResponse := handlers.BootstrapServer(getUserRequest, routeHandlers)
	responseBody = handlers.DecodeResponse(t, getUserResponse)
	assert.Equal(t, "user successfully retrieved", responseBody["message"])
}

func TestGetUser400(t *testing.T) {
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

	updateUserResponse := handlers.BootstrapServer(updateUserRequest, routeHandlers)
	responseBody := handlers.DecodeResponse(t, updateUserResponse)
	assert.Equal(t, "user successfully updated", responseBody["message"])

	getUserRequest := handlers.MakeTestRequest(t, fmt.Sprintf("/auth/getUser/%s", "wronguser.ID"), map[string]interface{}{}, "GET", &token)

	getUserResponse := handlers.BootstrapServer(getUserRequest, routeHandlers)
	responseBody = handlers.DecodeResponse(t, getUserResponse)
	assert.Equal(t, "not authorized", responseBody["message"])
}

func TestGetUser400NoID(t *testing.T) {
	code, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(user)

	verifyUser := handlers.MakeTestRequest(t, "/verifyOTP", map[string]interface{}{
		"id":  user.ID,
		"otp": code,
	}, "POST", nil)

	verifyResponse := handlers.BootstrapServer(verifyUser, routeHandlers)
	verifyResponseBody := handlers.DecodeResponse(t, verifyResponse)
	token := verifyResponseBody["token"]

	getUserRequest := handlers.MakeTestRequest(t, fmt.Sprintf("/getUser/%s", ""), map[string]interface{}{}, "GET", &token)

	getUserResponse := handlers.BootstrapServer(getUserRequest, routeHandlers)

	assert.Equal(t, 404, getUserResponse.Code)
}
