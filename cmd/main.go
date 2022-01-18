package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/gwuah/tinderclone/internal/middlewares"
)

func main() {
	// postgres.Init()
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.GET("/healthcheck", handlers.HealthGet)
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
