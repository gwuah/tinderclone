package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/models"
)

func (h *Handler) UpdateAccount(c *gin.Context) {

	type UpdateAccountRequest struct {
		ID        string `json:"id" binding:"required"`
		FirstName string `json:"first_name" binding:"required"`
		DOB       string `json:"dob" binding:"required"`
		Location  string `json:"location" binding:"required"`
	}

	var u UpdateAccountRequest
	var user models.User

	if err := c.BindJSON(&u); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	user.ID = u.ID
	user.DOB = lib.GetDob(u.DOB)
	user.Location = u.Location
	user.FirstName = u.FirstName

	err := h.repo.UserRepo.Update(&user)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update user",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully updated",
		"data"	 : user,
	})
}