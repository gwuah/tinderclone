package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/core/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func VerifyCodePost(db *gorm.DB, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u models.User

		if c.BindJSON(&u) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to create user. check documentation: https://github.com/gwuah/tinderclone#readme",
			})
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(u.OTP), []byte(password))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "code does not match OTP. try again?",
			})
		}

		c.JSON(http.StatusAccepted, gin.H{
			"message": "OTP code verified",
			"data":    u,
		})

	}

}
