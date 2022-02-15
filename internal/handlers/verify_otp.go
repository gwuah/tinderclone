package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/lib"
	"golang.org/x/crypto/bcrypt"
)

type VerifyOTPRequest struct {
	ID  string `json:"id"`
	OTP string `json:"otp"`
}

func (h *Handler) VerifyOTP(c *gin.Context) {
	var requestData VerifyOTPRequest
	var message *lib.Message

	if c.BindJSON(&requestData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}
	if requestData.ID == "" || requestData.OTP == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "must provide an OTP and an ID. fields cannot be left empty",
		})
		return
	}

	user, err := h.repo.UserRepo.FindUserByID(requestData.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no user found with that id"})
		return
	}

	if time.Now().After(user.OTPCreatedAt) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "otp has expired. regenerate a new one"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.OTP), []byte(requestData.OTP)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to verify otp"})
		return
	}

	token, err := lib.GenerateJWTToken(*user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "failed to generate jwt token"})
		return
	}
	
	sms := lib.NewSMS(os.Getenv("SMS_API_KEY"))
	message.To = user.PhoneNumber
	_, err = sms.SendTextMessage(*message)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to send OTP using SMS"})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "otp code verified",
		"data":    requestData,
		"token":   token,
	})

}
