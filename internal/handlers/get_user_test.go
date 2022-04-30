package handlers_test

import (
	"fmt"
	"testing"

	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/stretchr/testify/assert"
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

	getUserRequest := handlers.MakeTestRequest(t, fmt.Sprintf("/auth/getUser/%s", user.ID), map[string]interface{}{}, "GET", &token)

	getUserResponse := handlers.BootstrapServer(getUserRequest, routeHandlers)
	responseBody := handlers.DecodeResponse(t, getUserResponse)
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

	getUserRequest := handlers.MakeTestRequest(t, fmt.Sprintf("/auth/getUser/%s", "wronguser.ID"), map[string]interface{}{}, "GET", &token)

	getUserResponse := handlers.BootstrapServer(getUserRequest, routeHandlers)
	responseBody := handlers.DecodeResponse(t, getUserResponse)
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
