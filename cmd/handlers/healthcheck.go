package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthGet() gin.HandlerFunc {
	// c.JSON(http.StatusOK, gin.H{"msg": "what do you have to ask Joe Biren?"} or )
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}

}
