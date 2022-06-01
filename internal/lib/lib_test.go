package lib

import (
	//"os"
	"github.com/stretchr/testify/assert"
	"testing"
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
