package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

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

	httpEngine := gin.Default()

	httpEngine.Use(middlewares.Cors())
	httpEngine.GET("/", handlers.HealthGet)
	httpEngine.GET("/healthCheck", handlers.HealthGet)
	httpEngine.POST("/createAccount", handlers.CreateAccountPost(db))

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: httpEngine,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Println("failed To ShutDown Server", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server Closed After Interruption")
		} else {
			log.Println("Unexpected Server Shutdown. err:", err)
		}
	}
}
