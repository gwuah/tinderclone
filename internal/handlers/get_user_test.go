package handlers_test

import (
	"fmt"
	"testing"

	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestGetUser200(t *testing.T) {
	// TODO: happy test case. jwt and user id in path match
	code, _, user := handlers.CreateTestUser(t)
	handlers.SeedDB(&user)

	// req1 := handlers.MakeTestRequest(t, "/verifyOTP", map[string]interface{}{
	// 	"id":  user.ID,
	// 	"otp": code,
	// })

	// response := handlers.BootstrapServer(req1, routeHandlers)
	// responseBody := handlers.DecodeResponse(t, response)
	token := responseBody["token"]
	// assert.NoError(t, "otp code verified", responseBody["message"])

	req := handlers.MakeTestRequest(t, fmt.Sprintf("/retrieveUser/:%s", user.ID), map[string]interface{}{
		"token": token,
	})
	response = handlers.BootstrapServer(req, routeHandlers)
	responseBody = handlers.DecodeResponse(t, response)
	assert.Equal(t, "user succesfully retrieved", responseBody["message"])
}

func TestRetrieveUser400(t *testing.T) {
	// TODO: unhappy test case. wrong id
}

func TestRetrieveUser400NoID(t *testing.T) {
	// TODO: No id passed.
}
