package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	var u models.User

	if c.BindJSON(&u) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	results := h.db.Where("phone_number = ?", u.PhoneNumber).Find(&u)
	if results.Error != nil {
		log.Println(results.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no user found with that phone_number."})
		return
	}

	if results.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists."})
		return
	}

	code, err := lib.GenerateOTP()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create OTP"})
		return
	}

	hashedCode, err := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to hash OTP"})
		return
	}

	u.OTP = string(hashedCode)

	err = h.db.Create(&u).Error
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
