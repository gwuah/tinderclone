package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// add otp_created at field to users and store when an otp is generated
// reference that for validity check
// write integration tests for verify otp
// seperate readme to endpoint
type VerifyOTPRequest struct {
	ID  uint   `json:"id"`
	OTP string `json:"otp"`
}

func (h *Handler) VerifyOTP(c *gin.Context) {
	var requestData VerifyOTPRequest
	var u models.User

	if c.BindJSON(&requestData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	results := h.db.Where("id = ?", requestData.ID).Find(&u)
	if results.Error != nil {
		log.Println(results.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no user found with that id."})
		return
	}

	// change
	if u.OTPCreatedAt.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Expired OTP. Generate a new OTP."})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.OTP), []byte(requestData.OTP))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "failed to validate user OTP."})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "OTP code verified."})

}
