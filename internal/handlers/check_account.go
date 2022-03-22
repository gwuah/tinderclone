package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestProfileScore struct {
	ID string `json:"id"`
}
type ProfileScores struct {
	FirstName    int `json:"firstname"`
	LastName     int `json:"lastname"`
	Location     int `json:"location"`
	Bio          int `json:"bio"`
	Gender       int `json:"gender"`
	DOB          int `json:"dob"`
	Interests    int `json:"interests"`
	ProfilePhoto int `json:"profile_photo"`
}

func (h *Handler) CheckUser(c *gin.Context) {
	var requestData requestProfileScore
	var score ProfileScores

	if c.BindJSON(&requestData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	user, err := h.repo.UserRepo.FindUserByID(requestData.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "user does not exist"})
		return
	}

	if user.FirstName != "" {
		score.FirstName = 5
	}
	if !user.DOB.IsZero() {
		score.DOB = 15
	}
	if user.Location != "" {
		score.Location = 15
	}
	// if user.LastName != "" {
	// 	score.LastName = 5
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

	//TODO: profile photo = 25

	c.JSON(http.StatusOK, gin.H{
		"data": score,
	})

}
