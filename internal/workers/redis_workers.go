package workers

import (
	"encoding/json"
	"fmt"

	"github.com/bgentry/que-go"
	"github.com/go-redis/redis"
	"github.com/gwuah/tinderclone/internal/queue"
)

type AddToInterestBucketWorker struct {
	RedisClient *redis.Client
}

type AddToInterestBucketPayload struct {
	Interests []string
	ID        string
}

type RemoveFromInterestBucketWorker struct {
	RedisClient *redis.Client
}

type RemoveFromInterestBucketPayload struct {
	Interests []string
	ID        string
}

func NewAddToInterestBucketWorker(redisClient *redis.Client) *AddToInterestBucketWorker {
	return &AddToInterestBucketWorker{
		RedisClient: redisClient,
	}
}

func NewRemoveFromInterestBucketWorker(redisClient *redis.Client) *RemoveFromInterestBucketWorker {
	return &RemoveFromInterestBucketWorker{
		RedisClient: redisClient,
	}
}

func (r *AddToInterestBucketWorker) AddUserToEachInterestBucket(interests []string, id string) error {
	pipe := r.RedisClient.TxPipeline()
	for _, interest := range interests {
		if err := pipe.SAdd(interest, id).Err(); err != nil {
			return err
		}
	}
	_, err := pipe.Exec()
	if err != nil {
		return fmt.Errorf("failed to complete Add transaction. err= %s", err)
	}
	return nil
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

func (r *AddToInterestBucketWorker) Identifier() queue.Job {
	return ADD_TO_INTEREST_BUCKETS
}

func (r *RemoveFromInterestBucketWorker) Identifier() queue.Job {
	return REMOVE_FROM_INTEREST_BUCKETS
}

func (r *AddToInterestBucketWorker) Worker() que.WorkFunc {
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
