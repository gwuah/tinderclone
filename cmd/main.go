package main

import (
	"log"
	"os"

	"github.com/gwuah/tinderclone/internal/config"
	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/postgres"
	"github.com/gwuah/tinderclone/internal/queue"
	"github.com/gwuah/tinderclone/internal/repository"
	"github.com/gwuah/tinderclone/internal/workers"

	"github.com/gwuah/tinderclone/internal/server"
)

func main() {
	err := config.LoadNormalConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.Init()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db)

	sms, err := lib.NewSMS("Tinder Clone", os.Getenv("SMS_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	q, err := queue.New()
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.New(repo, sms, q)
	server := server.New(handler)

	w := q.RegisterJobs([]queue.JobWorker{
		workers.NewSMSWorker(sms),
	})
	go w.Start()

	server.Start()

}
