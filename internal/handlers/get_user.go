package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Scores struct {
	FirstName    int `json:"firstname" gorm:"-"`
	LastName     int `json:"lastname" gorm:"-"`
	Location     int `json:"location" gorm:"-"`
	Bio          int `json:"bio" gorm:"-"`
	Gender       int `json:"gender" gorm:"-"`
	DOB          int `json:"dob" gorm:"-"`
	Interests    int `json:"interests" gorm:"-"`
	ProfilePhoto int `json:"profile_photo" gorm:"-"`
}

func (h *Handler) GetUser(c *gin.Context) {
	var score Scores

	authorizedUserID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "request failed. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
	}

	if c.Param("id") != authorizedUserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
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
		"score":   score,
	})
}
