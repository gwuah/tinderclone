package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/core/config"
	"github.com/gwuah/tinderclone/internal/core/postgres"
	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/gwuah/tinderclone/internal/middlewares"
)

func main() {

	err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	_, err = postgres.Init()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(middlewares.Cors())
	r.GET("/", handlers.HealthGet)
	r.GET("/healthcheck", handlers.HealthGet)
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
