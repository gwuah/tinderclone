package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/lib"
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
		token, claims, err := lib.VerifyJWT(tokenString)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "token contains an invalid number of segments"})
			return
		}

		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user_id", claims.UserID)

	}
}
