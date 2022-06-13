package handlers

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/models"
	"github.com/gwuah/tinderclone/internal/workers"
	"github.com/thoas/go-funk"
)

type UpdateAccountRequest struct {
	ID           string          `json:"id" binding:"required"`
	FirstName    string          `json:"first_name"`
	LastName     string          `json:"last_name"`
	DOB          string          `json:"dob"`
	Location     models.Location `json:"location"`
	Bio          string          `json:"bio"`
	Gender       string          `json:"gender"`
	Interests    []string        `json:"interests"`
	ProfilePhoto string          `json:"profile_photo"`
}

func (h *Handler) UpdateAccount(c *gin.Context) {

	var u UpdateAccountRequest

	if err := c.BindJSON(&u); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	authorizedUserID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "request failed. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	if u.ID != authorizedUserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
		return
	}

	user := models.User{
		ID:           u.ID,
		DOB:          lib.GetDob(u.DOB),
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Bio:          u.Bio,
		Location:     u.Location,
		Gender:       u.Gender,
		Interests:    lib.SliceToString(u.Interests),
		ProfilePhoto: u.ProfilePhoto,
	}

	err := h.repo.UserRepo.UpdateUserByID(user.ID, &user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update user",
		})
		return
	}

	interests, err := h.repo.UserRepo.FindUserInterests(user.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to find interests",
		})
	}
	if len(interests) > 0 {
		if !reflect.DeepEqual(u.Interests, interests) {
			toAdd, toRemove := funk.Difference(u.Interests, interests)
			err = h.q.QueueJob(workers.ADD_TO_INTEREST_BUCKETS, workers.AddToInterestBucketPayload{
				Interests: toAdd,
				ID:        u.ID,
			})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to populate redis"})
				return
			}

			if err = h.q.QueueJob(workers.REMOVE_FROM_INTEREST_BUCKETS, workers.RemoveFromInterestBucketPayload{
				Interests: toRemove,
				ID:        u.ID,
			}); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to depopulate redis"})
				return
			}
		} else {
			// no change happened. skip or call redis sadd again.
			err = h.q.QueueJob(workers.ADD_TO_INTEREST_BUCKETS, workers.AddToInterestBucketPayload{
				Interests: u.Interests,
				ID:        u.ID,
			})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to populate redis"})
				return
			}
		}
	} else {
		err = h.q.QueueJob(workers.ADD_TO_INTEREST_BUCKETS, workers.AddToInterestBucketPayload{
			Interests: u.Interests,
			ID:        u.ID,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to populate redis"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully updated",
		"data":    user,
	})
}
