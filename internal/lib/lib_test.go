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

func TestSendSMS(t *testing.T) {
	sms, err := NewTermii("")
	assert.NoError(t, err)

	_, err = sms.SendTextMessage("233548669560", "Wo ho fin")

	assert.NoError(t, err)
}
