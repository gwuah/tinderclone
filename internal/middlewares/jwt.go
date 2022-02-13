package middlewares

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gwuah/tinderclone/internal/models"
)

var tokenKey = []byte("This is a secret key")

// TODO: generate more complex key

type JWTAuthDetails struct {
	jwt.StandardClaims
}

func CreateAccessToken(user models.User) (string, error) {
	expiresAt := time.Now().Add(time.Hour * 24 * 7).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTAuthDetails{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.ID,
			ExpiresAt: expiresAt,
		},
	})

	tokenString, err := token.SignedString(tokenKey)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}

func VerifyAccessToken(tokenString string) (uint, string, error) {
	var claims JWTAuthDetails
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tokenKey, nil
	})
	if err != nil {
		return 0, "", err
	}
	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	username := claims.Subject
	return 0, username, nil
}
