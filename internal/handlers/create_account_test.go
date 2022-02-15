package handlers_test

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gwuah/tinderclone/internal/config"
	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/postgres"
	"github.com/gwuah/tinderclone/internal/repository"
	"github.com/gwuah/tinderclone/internal/server"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := config.LoadTestConfig("../../.env.test")
	assert.NoError(&testing.T{}, err)

	db, err := postgres.Init()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db, lib.NewSMS(os.Getenv("SMS_API_KEY")))
	handler := handlers.New(repo)
	srv := server.New(handler)

	// defer srv.Stop()
	go srv.Start()

	os.Exit(m.Run())
}

func TestCreateAccountEndpoint(t *testing.T) {
	f := faker.New()

	req := map[string]interface{}{
		"phone_number": f.Numerify("+##############"),
	}

	resp, err := handlers.MakeRequest("createAccount", os.Getenv("PORT"), req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var m map[string]interface{}
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&m))

	defer resp.Body.Close()

	assert.Equal(t, "user succesfully created", m["message"])

}
