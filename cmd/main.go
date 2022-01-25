package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/config"
	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/gwuah/tinderclone/internal/middlewares"
	"github.com/gwuah/tinderclone/internal/postgres"
	"github.com/gwuah/tinderclone/internal/queue"
)

func main() {
	err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.Init()
	if err != nil {
		log.Fatal(err)
	}

	q, err := queue.New()
	if err != nil {
		log.Fatal(err)
	}

	workers := q.RegisterJobs([]queue.JobWorker{})
	go workers.Start()

	r := gin.Default()
	handler := handlers.New(db)
	r.Use(middlewares.Cors())
	r.GET("/healthCheck", handler.HealthCheck)
	r.POST("/createAccount", handler.CreateAccount)
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
