package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/cmd/handlers"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true

	r.GET("/healthcheck", handlers.HealthGet())

	r.Use(cors.New(config))
	r.Run(":6969")
}
