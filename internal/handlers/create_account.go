package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/models"
	"github.com/gwuah/tinderclone/internal/workers"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	var existingUser *models.User

	if c.BindJSON(existingUser) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	if existingUser.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "must provide a phone number. field cannot be left empty",
		})
		return
	}

	existingUser, rowsAffected, err := h.repo.UserRepo.FindUserByPhone(existingUser.PhoneNumber)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no user found with that phone number"})
		return
	}

	var sanitizedTermiiPhone string
	if string(existingUser.PhoneNumber[0]) == string("0") {
		sanitizedTermiiPhone = existingUser.CountryCode + strings.TrimPrefix(existingUser.PhoneNumber, "0")
	} else {
		sanitizedTermiiPhone = existingUser.CountryCode + existingUser.PhoneNumber
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

	existingUser.OTP = string(hashedCode)
	existingUser.OTPCreatedAt = lib.GenerateOTPExpiryDate()

	if rowsAffected > 0 {
		err := h.repo.UserRepo.Update(existingUser)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update otp for existing user"})
			return
		}

		err = h.q.QueueJob(workers.SEND_SMS, workers.SMSPayload{
			To:  sanitizedTermiiPhone,
			Sms: generateOTPMessage(code),
		})

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to queue sms otp"})
			return
		}

		existingUser.Sanitize()
		c.JSON(http.StatusOK, gin.H{
			"message": "user already existsjf",
			"data":    existingUser,
		})

		return
	}

	if err = h.repo.UserRepo.CreateUser(existingUser); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user"})
		return
	}

	err = h.q.QueueJob(workers.SEND_SMS, workers.SMSPayload{
		To:  sanitizedTermiiPhone,
		Sms: generateOTPMessage(code),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to queue sms otp"})
		return
	}

	existingUser.Sanitize()
	c.JSON(http.StatusCreated, gin.H{
		"message": "user succesfully created",
		"data":    existingUser,
	})
}

func generateOTPMessage(otp string) string {
	return fmt.Sprintf("Your tinderclone otp code is - %s", otp)
}
