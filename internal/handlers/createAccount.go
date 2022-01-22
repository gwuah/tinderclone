package handlers

import (
	"fmt"
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
				"message": "failed to create user. check documentation: https://github.com/gwuah/tinderclone#readme",
			})
			return
		}

		results := db.Where("phone_number = ?", u.PhoneNumber).Find(&u); if results.Error != nil{
			fmt.Println(results.Error)
		}

		if results.RowsAffected > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists"})
			return 
		}

		db := db.Create(&u); if db.Error != nil {
				fmt.Println(db.Error)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user"})
				return
			}

		
		c.JSON(http.StatusCreated, gin.H{
			"message": "user succesfully created",
			"data": u,
		})
	
	}}


