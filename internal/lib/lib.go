package lib

import (
	"crypto/rand"
	"net/url"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func GetDob(date string) time.Time {
	dateOfBirth, _ := time.Parse("02/01/2006", date)
	return dateOfBirth
}

func IsValidUrl(str string) bool {
	u, err := url.Parse(str)
	if err == nil && u.Scheme != "" && u.Host != "" {
		return true
	}
	return false
}
