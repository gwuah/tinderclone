package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/core/models"
	"github.com/gwuah/tinderclone/internal"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func CreateAccountPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u models.User

		if c.BindJSON(&u) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to create user. check documentation: https://github.com/gwuah/tinderclone#readme",
			})
			return
		}

		results := db.Where("phone_number = ?", u.PhoneNumber).Find(&u)
		if results.Error != nil {
			log.Println(results.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user."})
			return
		}

		if results.RowsAffected > 0 {

			c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists."})
			return
		}

		code, err := internal.GenerateOTP()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create OTP"})
		}

		hashedCode, err := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to hash OTP"})
		}

		u.OTP = string(hashedCode)

		err = db.Create(&u).Error
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user."})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "user succesfully created.",
			"data":    u,
		})

	}
}
