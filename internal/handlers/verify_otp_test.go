//go:build integration
// +build integration

package handlers_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/gwuah/tinderclone/internal/handlers"

	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

// find out what that testmain in first test does in create account

func TestVerifyOTPEndpoint(t *testing.T) {

	f := faker.New()

	createReq := map[string]interface{}{
		"phone_number": f.Numerify("+##############"),
	}

	createResp, createErr := handlers.MakeRequest("createAccount", os.Getenv("PORT"), createReq)
	assert.NoError(t, createErr)

	assert.Equal(t, http.StatusCreated, createResp.StatusCode)

	var m map[string]interface{}
	assert.NoError(t, json.NewDecoder(createResp.Body).Decode(&m))

	defer createResp.Body.Close()

	assert.Equal(t, "user succesfully created.", m["message"])

	verifyReq := map[string]interface{}{
		"id":  f.UInt8(),
		"otp": f.Numerify("#####"),
	}

	verifyResp, verifyErr := handlers.MakeRequest("verifyOTP", os.Getenv("PORT"), verifyReq)
	assert.NoError(t, verifyErr)

	assert.Equal(t, http.StatusInternalServerError, verifyResp.StatusCode)

	var o map[string]interface{}
	assert.NoError(t, json.NewDecoder(verifyResp.Body).Decode(&o))

	defer verifyResp.Body.Close()

	assert.Equal(t, "no user found with that id.", m["message"])

}
