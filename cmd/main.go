package main

import (
	"log"
	"os"

	"github.com/gwuah/tinderclone/internal/config"
	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/postgres"
	"github.com/gwuah/tinderclone/internal/queue"
	redistest "github.com/gwuah/tinderclone/internal/redis"
	"github.com/gwuah/tinderclone/internal/repository"
	"github.com/gwuah/tinderclone/internal/server"
	"github.com/gwuah/tinderclone/internal/workers"
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

	sms, err := lib.NewTermii(os.Getenv("SMS_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     os.Getenv("REDIS_URL"),
	// 	Password: os.Getenv("REDIS_PASSWORD"),
	// 	DB:       0,
	// })

	rdb, err := redistest.Init()
	if err != nil {
		log.Fatal(err)
	}

	q, err := queue.New()
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.New(repo, sms, q, rdb)
	server := server.New(handler)

	w := q.RegisterJobs([]queue.JobWorker{
		workers.NewSMSWorker(sms),
		workers.NewAddToInterestBucketWorker(rdb),
		workers.NewRemoveFromInterestBucketWorker(rdb),
	})
	go w.Start()

	server.Start()
}
