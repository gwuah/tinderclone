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
	tests := map[string]struct {
		input    string
		expected []string
	}{
		"no chars":        {input: "", expected: []string{""}},
		"simple":          {input: "a,b,c", expected: []string{"a", "b", "c"}},
		"trailing commas": {input: ",kwaku,richie,griff,dibri,", expected: []string{"kwaku", "richie", "griff", "dibri"}},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got := StringToSlice(testCase.input)
			if !assert.Equal(t, testCase.expected, got) {
				t.Fatalf("expected: %#v, got: %#v", testCase.expected, got)
			}
		})
	}

}
