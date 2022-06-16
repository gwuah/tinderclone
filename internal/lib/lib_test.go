package lib

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOTP(t *testing.T) {
	otp, err := GenerateOTP()
	assert.NoError(t, err)
	assert.Equal(t, 5, len(otp))
}

func TestSendSMS(t *testing.T) {
	sms, err := NewTermii(os.Getenv("SMS_API_KEY"))
	assert.NoError(t, err)

	_, err = sms.SendTextMessage("", "test")
	assert.NoError(t, err)
}

func TestStringToSlice(t *testing.T) {
	var outputSlice []string
	testString := []string{"", "a,b,c", "kwaku,richie,griff,dibri"}
	for _, v := range testString {
		testSlice := StringToSlice(v)
		outputSlice = append(outputSlice, testSlice...)
	}

	expectedOutput := []string{"", "a", "b", "c", "kwaku", "richie", "griff", "dibri"}
	assert.Equal(t, expectedOutput, outputSlice)
}
