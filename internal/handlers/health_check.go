package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, "tinderclone-api")
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
