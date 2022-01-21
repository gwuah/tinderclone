package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/core/models"
	"gorm.io/gorm"
)

func CreateAccountPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.User

		if c.BindJSON(&data) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error. no phone number received",
			})
			return
		}

		user := models.User{PhoneNumber: data.PhoneNumber}
		if err := db.Create(&user).Error; err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "phone number succesfully added",
			"data": gin.H{
				"id":     data.ID,
				"digits": data.PhoneNumber,
			},
		})

	}

}
