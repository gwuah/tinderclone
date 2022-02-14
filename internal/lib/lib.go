package lib

import (
	"crypto/rand"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gwuah/tinderclone/internal/models"
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

type JWTAuthDetails struct {
	jwt.StandardClaims
}

func GenerateJWTToken(user models.User) (string, error) {
	expiresAt := time.Now().Add(time.Hour * 24 * 7).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTAuthDetails{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.PhoneNumber,
			ExpiresAt: expiresAt,
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTOKENKEY")))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}
