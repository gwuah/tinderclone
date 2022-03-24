package lib

import (
	//"os"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestGenerateOTP(t *testing.T) {
	otp, err := GenerateOTP()
	assert.NoError(t, err)
	assert.Equal(t, 5, len(otp))
}

// func TestSendSMS(t *testing.T) {
// 	sms, err := NewTermii(os.Getenv("SMS_API_KEY"))
// 	assert.NoError(t, err)

// 	_, err = sms.SendTextMessage("", "test")

// 	assert.NoError(t, err)
// }

func TestIsValidUrl(t *testing.T) {
	f := faker.New()

	valid := IsValidUrl(f.Internet().URL())
	assert.True(t, valid)
}
