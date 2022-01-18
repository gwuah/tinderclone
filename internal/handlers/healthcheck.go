package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthGet(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
