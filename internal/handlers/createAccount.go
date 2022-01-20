package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/core/models"
)

func CreateAccountPost(c *gin.Context) {
	var data models.User

	if c.BindJSON(&data) != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "error. no phone number received",
		})
		return
	}

	digits := data.PhoneNumber
	id := data.ID

	c.JSON(http.StatusAccepted, gin.H{
		"message": "phone number succesfully added",
		"data": gin.H{
			"id":     id,
			"digits": digits,
		},
	})

}
