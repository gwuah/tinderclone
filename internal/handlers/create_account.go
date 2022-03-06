package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/models"
	"github.com/gwuah/tinderclone/internal/workers"
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

	code, err := lib.GenerateOTP()

	if rowsAffected > 0 {
		_, err = h.sms.SendTextMessage(u.PhoneNumber, generateOTPMessage(code))
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists, otp sent to user"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create otp"})
		return
	}

	hashedCode, err := lib.HashOTP(code)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to hash otp"})
		return
	}

	u.OTP = string(hashedCode)
	u.OTPCreatedAt = lib.GenerateOTPExpiryDate()

	if err = h.repo.UserRepo.CreateUser(&u); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user"})
		return
	}

	err = h.q.QueueJob(workers.SEND_SMS, workers.SMSPayload{
		To:  u.PhoneNumber,
		Sms: generateOTPMessage(code),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to que sms otp"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user succesfully created",
		"data":    u,
	})
}

func generateOTPMessage(otp string) string {
	return fmt.Sprintf("Your otp code is - %s", otp)
}
