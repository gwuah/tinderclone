package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/lib"
	"log"
	"net/http"
	"strings"
)

//TODO: Update documetation when authenticated routes exist
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no token found in authorization header"})
			return
		}
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid format in authorization header"})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := lib.VerifyAccessToken(tokenString)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "token contains an invalid number of segments"})
			return
		}

		_, ok := token.Claims.(*lib.JWTAuthDetails)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid claims used for token"})
			return
		}

		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
