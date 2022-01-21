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
		var u models.User

		if c.BindJSON(&u) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error. no phone number received",
			})
			return
		}

		if err := db.Create(&u).Error; err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "phone number succesfully added",
			"data": u,
		})

	}

}
