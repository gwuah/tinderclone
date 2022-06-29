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

func TestFindDifferenceBetweenInterests(t *testing.T) {
	type Case struct {
		name       string
		input      [][]string
		difference []string
	}

	cases := []Case{
		Case{
			name:       "simple",
			input:      [][]string{[]string{"yaw", "asare"}, []string{"asare"}},
			difference: []string{"yaw"},
		},
		Case{
			name:       "simple2",
			input:      [][]string{[]string{"asare"}, []string{"yaw", "asare"}},
			difference: []string{"yaw"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := FindDifferenceBetweenInterests(tc.input[0], tc.input[1])
			assert.Equal(t, tc.difference, got)
		})
	}

}

func TestIntersection(t *testing.T) {
	type Case struct {
		name       string
		input      [][]string
		difference []string
	}

	cases := []Case{
		Case{
			name:       "simple",
			input:      [][]string{[]string{"yaw", "asare"}, []string{"asare"}},
			difference: []string{"asare"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Intersection(tc.input[0], tc.input[1])
			assert.Equal(t, tc.difference, got)
		})
	}

}

func TestComplement(t *testing.T) {
	type Case struct {
		name              string
		previousInterests []string
		currentInterests  []string

		toAdd    []string
		toRemove []string
	}

	cases := []Case{
		Case{
			name: "complex",

			previousInterests: []string{"hiking"},
			currentInterests:  []string{"hiking", "singing"},

			toAdd:    []string{"singing"},
			toRemove: []string{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			unchangedInterests := Intersection(tc.previousInterests, tc.currentInterests)
			toRemove := Complement(unchangedInterests, tc.previousInterests)
			toAdd := Complement(unchangedInterests, tc.currentInterests)

			assert.Equal(t, tc.toAdd, toAdd)
			assert.Equal(t, tc.toRemove, toRemove)

		})
	}

}
