package middlewares

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt"
// 	"github.com/gwuah/tinderclone/internal/lib"
// 	"log"
// 	"net/http"
// )

// func AuthorizeJWT() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		const BEARER_SCHEMA = "Bearer"
// 		authHeader := c.GetHeader("Authorization")
// 		tokenString := authHeader[len(BEARER_SCHEMA):]
// 		token, err := lib.VerifyAccessToken(tokenString)
// 		if err != nil {
// 			log.Println(err)
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "unable to parse token"})
// 		}

// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		log.Println(claims)
// 		if !ok {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid claims"})
// 		}


// 		if !token.Valid {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}

// 	}
// }
