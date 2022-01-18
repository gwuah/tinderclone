package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, "tinderclone-api")
}
func HealthGet(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
