package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

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

	handler := handlers.New(db)
	httpEngine := gin.Default()

	httpEngine.Use(middlewares.Cors())
	httpEngine.GET("/healthCheck", handler.HealthCheck)
	httpEngine.POST("/createAccount", handler.CreateAccount)
	httpEngine.POST("/verifyOTP", handler.VerifyOTP)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: httpEngine,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Println("failed to shutdown server", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("server closed after interruption")
		} else {
			log.Println("unexpected server shutdown. err:", err)
		}
	}
}
