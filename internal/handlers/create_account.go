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

	if u.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "must provide a phone number. field cannot be left empty",
		})
		return
	}

	_, rowsAffected, err := h.repo.UserRepo.FindUserByPhone(u.PhoneNumber)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no user found with that phone number"})
		return
	}

	if rowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists"})
		return
	}

	code, err := lib.GenerateOTP()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create otp"})
		return
	}

	hashedCode, err := lib.HashOTP(code)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to hash otp"})
		return
	}

	// TODO: JSON object no longer returns ID since UUID change. u.ID =
	u.OTP = string(hashedCode)
	u.OTPCreatedAt = lib.GenerateOTPExpiryDate()

	if err = h.repo.UserRepo.CreateUser(&u); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user succesfully created",
		"data":    u,
	})
}
