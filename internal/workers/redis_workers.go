package workers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bgentry/que-go"
	"github.com/go-redis/redis"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/queue"
)

type RedisWorker struct {
	RedisClient *redis.Client
}

type RedisPayload struct {
	StringOfInterests string
	ID                string
}

func NewRedisWorker(redisClient *redis.Client) *RedisWorker {
	return &RedisWorker{
		RedisClient: redisClient,
	}
}

func (r *RedisWorker) UpdateBucketWithInterests(StringOfInterests string, id string) error {
	sliceOfInterests := lib.StringToSlice(StringOfInterests)
	for _, element := range sliceOfInterests {
		x := r.RedisClient.SAdd(element, id)
		if err := x.Err(); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (r *RedisWorker) Identifier() queue.Job {
	return UPDATE_REDIS_BUCKET
}

func (r *RedisWorker) Worker() que.WorkFunc {
	return func(j *que.Job) error {
		var req RedisPayload
		if err := json.Unmarshal(j.Args, &req); err != nil {
			return fmt.Errorf("unmarshal job failed. args= %s | err= %w", string(j.Args), err)
		}

		err := r.UpdateBucketWithInterests(req.StringOfInterests, req.ID)
		if err != nil {
			return fmt.Errorf("failed to populate redis bucket, error: \n %w", err)
		}
		
		return nil
	}
}
