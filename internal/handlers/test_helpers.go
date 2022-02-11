package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/models"
	"github.com/gwuah/tinderclone/internal/postgres"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

	hasedCode, err := lib.HashOTP(code)
	assert.NoError(t, err)

	testUser := models.User{
		ID:           NewUUID(),
		PhoneNumber:  f.Numerify("+##############"),
		OTP:          string(hasedCode),
		OTPCreatedAt: lib.GenerateOTPExpiryDate(),
	}

	return code, string(hasedCode), &testUser
}
