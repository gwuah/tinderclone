package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gwuah/tinderclone/internal/lib"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		var claims lib.JWTAuthDetails
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return os.Getenv("JWTOKENKEY"), nil
		})
		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println(claims)
			log.Println("valid token")
		} else {
			log.Println(err)
			log.Println("invalid token")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func VerifyAccessToken(tokenString string) (string, error) {
	var claims lib.JWTAuthDetails
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return os.Getenv("JWTOKENKEY"), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}

	username := claims.Subject
	return username, nil
}
