package queue

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/bgentry/que-go"
	"github.com/gwuah/tinderclone/internal/postgres"
	"github.com/jackc/pgx"
)

type Job string

func (j Job) String() string {
	return string(j)
}

type (
	JobWorker interface {
		Identifier() Job
		Worker() que.WorkFunc
	}
	QueImpl interface {
		Close()
		RegisterJobs(jobList []JobWorker) *que.WorkerPool
		QueueJob(jobType Job, payload interface{}) error
		QueueFutureJob(jobType Job, payload interface{}, time ...time.Time) error
	}
	Que struct {
		dbURI    string
		client   *que.Client
		connPool *pgx.ConnPool
	}
)

func getPgxPool(dbUri string) (*pgx.ConnPool, error) {
	pgxcfg, err := pgx.ParseURI(dbUri)
	if err != nil {
		return nil, err
	}
	pgxpool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgxcfg,
		// Our hosted postgres instance only allows us to make 20 connections to the database.
		// Mind you this is a different connection than the one inside postgres.Init()
		// See here - https://data.heroku.com/datastores/75c50280-a8be-4ab8-9be6-e2ce4ed24839
		MaxConnections: 15,
		AfterConnect:   que.PrepareStatements,
	})
	if err != nil {
		return nil, err
	}
	return pgxpool, nil
}

func New() (*Que, error) {
	q := &Que{dbURI: postgres.ConstructDatabaseURI()}
	pgxpool, err := getPgxPool(q.dbURI)
	if err != nil {
		return nil, err
	}
	q.connPool = pgxpool
	q.client = que.NewClient(pgxpool)
	return q, nil
}

func (q *Que) Close() {
	log.Println("shutting down queue")
	q.connPool.Close()
}

func (q *Que) RegisterJobs(jobList []JobWorker) *que.WorkerPool {
	wm := que.WorkMap{}
	for _, j := range jobList {
		wm[j.Identifier().String()] = j.Worker()
	}
	return que.NewWorkerPool(q.client, wm, 10)
}

func (q *Que) QueueJob(jobType Job, payload interface{}) error {
	enc, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	j := que.Job{Type: jobType.String(), Args: enc}
	err = q.client.Enqueue(&j)
	if err != nil {
		return fmt.Errorf("failed to queue job. err: %w", err)
	}
	return nil
}

func (q *Que) QueueFutureJob(jobType Job, payload interface{}, times ...time.Time) error {
	enc, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	for _, time := range times {
		j := que.Job{Type: jobType.String(), Args: enc, RunAt: time}
		err = q.client.Enqueue(&j)
		if err != nil {
			return fmt.Errorf("failed to queue job. err: %w", err)
		}
	}

	return nil
}
