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

	return a
}

func getKeys(m map[string]bool) []string {
	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func Intersection(a, b []string) []string {
	intersection := map[string]bool{}

	a_map := map[string]bool{}
	b_map := map[string]bool{}

	for _, v := range a {
		a_map[v] = true
	}

	for _, v := range b {
		b_map[v] = true
	}

	if len(a_map) > len(b_map) {
		a_map, b_map = b_map, a_map
	}

	for k := range a_map {
		if b_map[k] {
			intersection[k] = true
		}
	}

	return getKeys(intersection)
}

func Complement(intersection, values []string) []string {
	v_map := map[string]bool{}
	for _, v := range values {
		v_map[v] = true
	}

	for _, v := range intersection {
		delete(v_map, v)
	}

	return getKeys(v_map)
}
