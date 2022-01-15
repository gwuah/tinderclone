package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}

}
