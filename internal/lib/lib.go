package lib

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

const otpChars = "1234567890"

func GenerateOTPExpiryDate() time.Time {
	expiry := time.Now().Add(time.Minute * 3)
	return expiry
}

func HashOTP(code string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
}

func GenerateOTP() (string, error) {
	buffer := make([]byte, 5)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpChars5 := len(otpChars)
	for i := 0; i < 5; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpChars5]
	}

	return string(buffer), nil
}

func MakeSMSRequest(endpoint string, requestBody interface{}) (*http.Response, error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
