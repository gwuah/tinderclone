package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// validate:"required"
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

	err := bcrypt.CompareHashAndPassword([]byte(u.OTP), []byte(requestData.OTP))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "failed to validate user OTP."})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "OTP code verified."})

}
