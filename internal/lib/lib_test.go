package lib

import (
	"testing"
)

func TestGenerateOTP(t *testing.T) {
	test, _ := GenerateOTP()
	if len(test) != 5 {
		t.Error("Test Failed: expected GenerateOTP function to return only 5 digits.")
	}

}
func TestOTPCharLength(t *testing.T) {
	if len(otpChars) != 10 {
		t.Error("Test Failed: expcted OTPChars to be of lenght 10.")
	}
}

// func TestOTPCharUnique(t *testing.T) {
// 	for _, num := range otpChars {
// 		if strings.Contains(otpChars, string(num)) {
// 			t.Error("Test Failed: all characters in otpChars should be unique.")
// 		}
// 	}
// }
