package lib

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gwuah/tinderclone/internal/models"
)

type JWTAuthDetails struct {
	UserID string
	jwt.StandardClaims
}

func GenerateJWTToken(user models.User) (string, error) {
	expiresAt := time.Now().Add(time.Hour * 24 * 7).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTAuthDetails{
		UserID: user.ID,
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

func VerifyAccessToken(tokenString string) (*jwt.Token, JWTAuthDetails, error) {
	var claims JWTAuthDetails
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWTOKENKEY")), nil
	})

	return token, claims, err
}
