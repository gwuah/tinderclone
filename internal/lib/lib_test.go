package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"

)

func TestGenerateOTP(t *testing.T) {
	otp, err := GenerateOTP()
	assert.NoError(t, err)
	assert.Equal(t, 5, len(otp))
}

// func TestSendSMS(t *testing.T) {
// 	sms, err := NewSMS("Tinder Clone", os.Getenv("SMS_API_KEY"))
// 	assert.NoError(t, err)

// 	_, err = sms.SendTextMessage("0205428210", "testing")

// 	assert.NoError(t, err)
// }
