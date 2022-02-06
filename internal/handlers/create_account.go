package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/models"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	var u models.User

	if c.BindJSON(&u) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	_, rowsAffected, err := h.repo.UserRepo.FindUserByPhone(u.PhoneNumber)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no user found with that phone number."})
		return
	}

	if rowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists."})
		return
	}

	code, err := lib.GenerateOTP()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create OTP"})
		return
	}

	hashedCode, err := lib.HashOTP(code)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to hash OTP"})
		return
	}

	u.OTP = string(hashedCode)
	u.OTPCreatedAt = lib.GenerateOTPExpiryDate()

	if err = h.repo.UserRepo.CreateUser(&u); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user succesfully created.",
		"data":    u,
	})
}
