package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/models"
)

type UpdateAccountRequest struct {
	ID           string `json:"id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	DOB          string `json:"dob" binding:"required"`
	Location     string `json:"location" binding:"required"`
	ProfilePhoto string `json:"profile_photo" binding:"required"`
}

func (h *Handler) UpdateAccount(c *gin.Context) {

	var u UpdateAccountRequest
	var user models.User

	if err := c.BindJSON(&u); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}
	validateUrl := lib.IsValidUrl(u.ProfilePhoto)
	if !validateUrl {
		log.Panicln("use a valid image url")
	}

	user = models.User{
		ID:           u.ID,
		DOB:          lib.GetDob(u.DOB),
		Location:     u.Location,
		FirstName:    u.FirstName,
		ProfilePhoto: u.ProfilePhoto,
	}

	err := h.repo.UserRepo.UpdateUserByID(&user)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully updated",
		"data":    user,
	})
}
