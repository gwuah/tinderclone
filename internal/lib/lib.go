package lib

import (
	"crypto/rand"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const otpChars = "1234567890"

func GenerateOTPExpiryDate() time.Time {
	expiry := time.Now().Add(time.Minute * 3)
	return expiry
}

func HashOTP(code string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
}

func GenerateOTP() (string, error) {
	buffer := make([]byte, 5)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpChars5 := len(otpChars)
	for i := 0; i < 5; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpChars5]
	}

	return string(buffer), nil
}

func GetDob(date string) time.Time {
	dateOfBirth, _ := time.Parse("02/01/2006", date)
	return dateOfBirth
}

func SliceToString(slice []string) string {
	stringifiedSlice := strings.Join(slice, ",")
	return stringifiedSlice
}

func StringToSlice(stringifiedSlice string) []string {
	temp := SanitizeString(stringifiedSlice)
	slice := strings.Split(temp, ",")
	return slice
}

func FindDifferenceBetweenInterests(a, b []string) []string {
	mapOfStrings := make(map[string]bool)
	for _, val := range b {
		mapOfStrings[val] = true
	}
	var difference []string
	for _, val := range a {
		if _, found := mapOfStrings[val]; !found {
			difference = append(difference, val)
		}
	}
	return difference
}

func EqualInterests(a, b []string) bool {
	return len(FindDifferenceBetweenInterests(a, b)) == 0
}

func SanitizeString(a string) (b string) {
	char := ","

	a = strings.TrimPrefix(a, char)
	a = strings.TrimSuffix(a, char)

	// if strings.HasPrefix(a, char) {
	// 	a = a[len(char):]
	// }

	// if strings.HasSuffix(a, char) {
	// 	a = a[:len(a)-len(char)]
	// }
	return a
}
