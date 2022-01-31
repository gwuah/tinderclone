package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gwuah/tinderclone/internal/config"
	"github.com/gwuah/tinderclone/internal/handlers"
	"github.com/gwuah/tinderclone/internal/middlewares"
	"github.com/gwuah/tinderclone/internal/postgres"
	"github.com/gwuah/tinderclone/internal/queue"
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

	q, err := queue.New()
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range os.Environ() {

		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%s: %s\n", pair[0], pair[1])
	}

	handler := handlers.New(db)
	server := server.New(handler, middlewares.Cors())

	workers := q.RegisterJobs([]queue.JobWorker{})
	go workers.Start()

	server.Start()

}
