package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/models"
)

func (h *Handler) GetUser(c *gin.Context) {
	var score models.Scores

	authorizedUserID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "request failed. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
	}

	if c.Param("id") != authorizedUserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
	}

	user, err := h.repo.UserRepo.FindUserByID(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no user found with that id"})
		return
	}

	if user.FirstName != "" {
		score.FirstName = 5
	}
	if !user.DOB.IsZero() {
		score.DOB = 15
	}
	// if user.LastName != "" {
	// 	score.LastName = 5
	// }
	// if user.Location != "" {
	// 	score.Location = 15
	// }
	// if user.Bio != "" {
	// 	score.Bio = 5
	// }
	// if user.Gender != "" {
	// 	score.Gender = 20
	// }
	// if user.Interests[0] != "" {
	// 	score.Interests = 10
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "user succesfully retrieved",
		"user":    user,
	})
}
