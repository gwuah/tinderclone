package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/core/postgres"
	"github.com/gwuah/tinderclone/handlers"
)

func main() {
	postgres.Init()

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true

	r.GET("/healthcheck", handlers.HealthGet)
	r.Use(cors.New(config))
	r.Run(":8000")
}
