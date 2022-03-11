package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type checkUserProfile struct {
	ID string `json:"id"`
}

func (h *Handler) CheckUser(c *gin.Context) {
	var requestData checkUserProfile
	var profileScore int

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
		_ = AddScore(profileScore, 5)
	}
	if user.LastName != "" {
		_ = AddScore(profileScore, 5)
	}
	if user.Location != "" {
		_ = AddScore(profileScore, 15)
	}
	if user.Bio != "" {
		_ = AddScore(profileScore, 5)
	}
	if user.Gender != "" {
		_ = AddScore(profileScore, 20)
	}
	if !user.DOB.IsZero() {
		_ = AddScore(profileScore, 15)
	}
	if user.Interests[0] != "" {
		_ = AddScore(profileScore, 10)
	}

	// profile photo = 25

	c.JSON(http.StatusOK, gin.H{
		"profile_score": profileScore,
		"data":          user,
	})

}

func AddScore(profileScore int, rubrik int) int {
	score := profileScore + rubrik
	return score
}
