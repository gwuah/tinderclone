package handlers_test

import (
	"testing"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/config"
	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/postgres"
	"github.com/gwuah/tinderclone/internal/queue"
	"github.com/gwuah/tinderclone/internal/repository"
	"github.com/gwuah/tinderclone/internal/server"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var routeHandlers *gin.Engine

func TestMain(m *testing.M) {
	err := config.LoadTestConfig("../../.env.test")
	assert.NoError(&testing.T{}, err)

	db, err := postgres.Init()
	if err != nil {
		log.Fatal(err)
	}

	sms, err := lib.NewTermii(os.Getenv("SMS_API_KEY"))
	assert.NoError(&testing.T{}, err)
	q, err := queue.New()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db)
	handler := handlers.New(repo, sms, q)
	srv := server.New(handler)
	routeHandlers = srv.SetupRoutes()
	os.Exit(m.Run())
}

func TestCreateAccountEndpoint(t *testing.T) {
	f := faker.New()

	req := handlers.MakeTestRequest(t, "/createAccount", map[string]interface{}{
		"phone_number": f.Numerify("+##############"),
	})

	response := handlers.BootstrapServer(req, routeHandlers)
	responseBody := handlers.DecodeResponse(t, response)

	assert.Equal(t, "user successfully created", responseBody["message"])

}

