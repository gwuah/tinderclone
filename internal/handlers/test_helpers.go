package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/models"
	"github.com/gwuah/tinderclone/internal/postgres"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var dbConnPool *gorm.DB

func NewUUID() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}

func MakeRequest(endpoint string, port string, requestBody interface{}) (*http.Response, error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:%s/%s", port, endpoint), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func SeedDB(r ...interface{}) *gorm.DB {
	if dbConnPool == nil {
		db, err := postgres.Init()
		if err != nil {
			log.Fatal(err)
		}
		dbConnPool = db
	}
	tx := dbConnPool.Begin()
	for _, m := range r {
		if err := tx.Create(m).Error; err != nil {
			tx.Rollback()
			log.Fatalf("[data insert failed] %v ", err)
		}
	}
	tx.Commit()
	return dbConnPool
}

func CreateTestUser(t *testing.T) (string, string, *models.User) {
	f := faker.New()

	code, err := lib.GenerateOTP()
	assert.NoError(t, err)

	hashedCode, err := lib.HashOTP(code)
	assert.NoError(t, err)

	testUser := models.User{
		ID:           NewUUID(),
		PhoneNumber:  f.Numerify("+##############"),
		OTP:          string(hashedCode),
		OTPCreatedAt: lib.GenerateOTPExpiryDate(),
	}

	return code, string(hashedCode), &testUser
}

func BootstrapServer(req *http.Request, routeHandlers *gin.Engine) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	routeHandlers.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func MakeTestRequest(t *testing.T, route string, body interface{}) *http.Request {
	reqBody, err := json.Marshal(body)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", route, bytes.NewReader(reqBody))
	assert.NoError(t, err)

	return req
}

func DecodeResponse(t *testing.T, response *httptest.ResponseRecorder) map[string]interface{} {
	var responseBody map[string]interface{}
	assert.NoError(t, json.Unmarshal(response.Body.Bytes(), &responseBody))
	return responseBody
}
