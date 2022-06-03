package workers

import (
	"encoding/json"
	"fmt"

	"github.com/bgentry/que-go"
	"github.com/go-redis/redis"
	"github.com/gwuah/tinderclone/internal/queue"
)

type AddToInterestBuckerWorker struct {
	RedisClient *redis.Client
}

type AddToInterestBucketPayload struct {
	Interests []string
	ID        string
}

func NewAddToInterestBuckerWorker(redisClient *redis.Client) *AddToInterestBuckerWorker {
	return &AddToInterestBuckerWorker{
		RedisClient: redisClient,
	}
}

func (r *AddToInterestBuckerWorker) AddUserToEachInterestBucket(interests []string, id string) error {
	for _, interest := range interests {
		if err := r.RedisClient.SAdd(interest, id).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (r *AddToInterestBuckerWorker) Identifier() queue.Job {
	return ADD_TO_INTEREST_BUCKETS
}

func (r *AddToInterestBuckerWorker) Worker() que.WorkFunc {
	return func(j *que.Job) error {
		var req AddToInterestBucketPayload
		if err := json.Unmarshal(j.Args, &req); err != nil {
			return fmt.Errorf("unmarshal job failed. args= %s | err= %w", string(j.Args), err)
		}

		err := r.AddUserToEachInterestBucket(req.Interests, req.ID)
		if err != nil {
			return fmt.Errorf("failed to populate redis bucket, error: \n %w", err)
		}
		return nil
	}
}
