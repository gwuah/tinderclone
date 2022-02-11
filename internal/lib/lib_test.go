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
