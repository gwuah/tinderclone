package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type VerifyOTPRequest struct {
	ID  string `json:"id"`
	OTP string `json:"otp"`
}

func (h *Handler) VerifyOTP(c *gin.Context) {
	var requestData VerifyOTPRequest

	if c.BindJSON(&requestData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	user, err := h.repo.UserRepo.FindUserByID(requestData.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no user found with that id"})
		return
	}

	if user.OTPCreatedAt.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "otp has expired. regenerate a new one"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.OTP), []byte(requestData.OTP)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to verify otp"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "otp code verified"})

}
