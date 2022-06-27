package lib

import (
	"log"
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
	tests := []struct {
		input    string
		expected []string
	}{
		{input: "", expected: []string{""}},
		{input: "a,b,c", expected: []string{"a", "b", "c"}},
		{input: "kwaku,richie,griff,dibri", expected: []string{"kwaku", "richie", "griff", "dibri"}},
	}

	for _, testCase := range tests {
		got := StringToSlice(testCase.input)
		log.Println(got)
		assert.Equal(t, testCase.expected, got)
	}

}
