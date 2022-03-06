package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/lib"
	"log"
	"net/http"
)

func (h *Handler) UpdateAccount(c *gin.Context) {

	type UpdateAccountRequest struct {
		ID        string `json:"id" binding:"required"`
		FirstName string `json:"first_name" binding:"required"`
		DOB       string `json:"dob" binding:"required"`
		Location  string `json:"location" binding:"required"`
	}

	var u UpdateAccountRequest

	if err := c.BindJSON(&u); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	user, err := h.repo.UserRepo.FindUserByID(u.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no user found with that id",
		})
	}

	user.DOB = lib.GetDob(u.DOB)
	user.Location = u.Location
	user.FirstName = u.FirstName

	err = h.repo.UserRepo.UpdateUser(user)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update user",
		})
	}
}
