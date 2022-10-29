package workers

import (
	"encoding/json"
	"fmt"

	"github.com/bgentry/que-go"
	"github.com/go-redis/redis"
	"github.com/gwuah/tinderclone/internal/queue"
)

type RemoveFromInterestBucketWorker struct {
	RedisClient *redis.Client
}

type RemoveFromInterestBucketPayload struct {
	Interests []string
	ID        string
}

func NewRemoveFromInterestBucketWorker(redisClient *redis.Client) *RemoveFromInterestBucketWorker {
	return &RemoveFromInterestBucketWorker{
		RedisClient: redisClient,
	}
}

func (r *RemoveFromInterestBucketWorker) RemoveUserFromEachInterestBucket(interests []string, id string) error {
	pipe := r.RedisClient.TxPipeline()

	for _, interest := range interests {
		if err := pipe.SRem(interest, id).Err(); err != nil {
			return err
		}
	}
	_, err := pipe.Exec()
	if err != nil {
		return fmt.Errorf("failed to complete Remove transaction. err= %s", err)
	}
	return nil
}

func (r *RemoveFromInterestBucketWorker) Identifier() queue.Job {
	return REMOVE_FROM_INTEREST_BUCKETS
}

func (r *RemoveFromInterestBucketWorker) Worker() que.WorkFunc {
	return func(j *que.Job) error {
		var req RemoveFromInterestBucketPayload
		if err := json.Unmarshal(j.Args, &req); err != nil {
			return fmt.Errorf("unmarshal job failed. args= %s | err= %w", string(j.Args), err)
		}

		err := r.RemoveUserFromEachInterestBucket(req.Interests, req.ID)
		if err != nil {
			return fmt.Errorf("failed to remove interest from redis bucket, error: \n %w", err)
		}
		return nil
	}
}
