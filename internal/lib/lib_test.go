package lib

import (
	"testing"
)

// type TestCase struct {
// 	expected int
// 	actual   int
// }
func TestGenerateOTP(t *testing.T) {
	test, _ := GenerateOTP()
	if len(test) != 5 {
		t.Error("Test Failed: expected GenerateOTP function to return only 5 digits.")
	}

}

// find out wtf a gwan
// func TestOTPCharLength(t *testing.T) {
// 	if len(otpChars) != 10 {
// 		t.Error("Test Failed: expcted OTPChars to be of lenght 10.")
// 	}
// }

// func TestOTPCharUnique(t *testing.T) {
// 	// exists := make(map[string]int)
// 	var exists1 string
// 	for i, letter := range otpChars {
// 		if strings.Contains(exists1, string(letter)) {
// 			i = i + 1
// 		}
// 		unique := exists1 + string(letter)
// 		return unique
// 	}
// }
