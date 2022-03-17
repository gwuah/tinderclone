package handlers

import (
	//"fmt"
	"log"
	"net/http"
	//"strings"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/models"
	//"github.com/gwuah/tinderclone/internal/workers"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	var newUser models.User

	err := c.BindJSON(&newUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	if newUser.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "must provide a phone number. field cannot be left empty",
		})
		return
	}

	existingUser, rowsAffected, err := h.repo.UserRepo.FindUserByPhone(newUser.PhoneNumber)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no user found with that phone number"})
		return
	}

	// var sanitizedTermiiPhone string
	// if string(newUser.PhoneNumber[0]) == "0" {
	// 	sanitizedTermiiPhone = newUser.CountryCode + strings.TrimPrefix(newUser.PhoneNumber, "0")
	// } else {
	// 	sanitizedTermiiPhone = newUser.CountryCode + newUser.PhoneNumber
	// }


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

	newUser.OTP = string(hashedCode)
	newUser.OTPCreatedAt = lib.GenerateOTPExpiryDate()
	newUser.RawOTP = code

	if rowsAffected > 0 {
		existingUser.RawOTP = code
		existingUser.OTP = string(hashedCode)
		existingUser.OTPCreatedAt = lib.GenerateOTPExpiryDate()
		err := h.repo.UserRepo.Update(existingUser)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update otp for existing user"})
			return
		}

		// err = h.q.QueueJob(workers.SEND_SMS, workers.SMSPayload{
		// 	To:  sanitizedTermiiPhone,
		// 	Sms: generateOTPMessage(code),
		// })

		// if err != nil {
		// 	log.Println(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to queue sms otp"})
		// 	return
		// }

		existingUser.Sanitize()
		c.JSON(http.StatusOK, gin.H{
			"message": "user already exists",
			"data":    existingUser,
		})

		return
	}

	if err = h.repo.UserRepo.CreateUser(&newUser); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user"})
		return
	}

	// err = h.q.QueueJob(workers.SEND_SMS, workers.SMSPayload{
	// 	To:  sanitizedTermiiPhone,
	// 	Sms: generateOTPMessage(code),
	// })
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to queue sms otp"})
	// 	return
	// }

	newUser.Sanitize()
	c.JSON(http.StatusCreated, gin.H{
		"message": "user successfully created",
		"data":    newUser,
	})
}

// func generateOTPMessage(otp string) string {
// 	return fmt.Sprintf("Your tinderclone otp code is - %s", otp)
// }
