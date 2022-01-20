package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/core/config"
	"github.com/gwuah/tinderclone/internal/core/postgres"
	"github.com/gwuah/tinderclone/internal/core/queue"
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

	q, err := queue.New()
	if err != nil {
		log.Fatal("failed to initialize queue. err", err)
	}

	workers := q.RegisterJobs([]queue.JobWorker{})
	go workers.Start()

	r := gin.Default()

	r.Use(middlewares.Cors())
	r.GET("/", handlers.HealthGet)
	r.GET("/healthcheck", handlers.HealthGet)
	r.POST("/createaccount", handlers.CreateAccountPost)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
