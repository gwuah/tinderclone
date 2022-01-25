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

	db, err := postgres.Init()
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
	// remove this or?
	r.GET("/healthCheck", handlers.HealthGet)
	r.POST("/createAccount", handlers.CreateAccountPost(db))
	r.POST("/verifyOTP", handlers.VerifyCodePost(db))
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
