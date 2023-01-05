package workers

import (
	"encoding/json"
	"fmt"

	"github.com/bgentry/que-go"
	"github.com/go-redis/redis"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/queue"
)

type UpdateUserWorker struct {
	RedisClient *redis.Client
	queue       *queue.Que
}

type UpdateUserWorkerPayload struct {
	UserID            string
	PreviousInterests []string
	CurrentInterests  []string
}

func (r *UpdateUserWorker) Identifier() queue.Job {
	return UPDATE_USER
}

func NewUpdateUserWorker(redisClient *redis.Client, queue *queue.Que) *UpdateUserWorker {
	return &UpdateUserWorker{
		RedisClient: redisClient,
		queue:       queue,
	}
}

func (r *UpdateUserWorker) Worker() que.WorkFunc {
	return func(j *que.Job) error {
		var req UpdateUserWorkerPayload
		if err := json.Unmarshal(j.Args, &req); err != nil {
			return fmt.Errorf("unmarshal job failed. args= %s | err= %w", string(j.Args), err)
		}

		if lib.EqualInterests(req.PreviousInterests, req.CurrentInterests) {
			return nil
		}

		if len(req.PreviousInterests) == 0 && len(req.CurrentInterests) > 0 {
			return r.queue.QueueJob(ADD_TO_INTEREST_BUCKETS, AddToInterestBucketPayload{
				Interests: req.CurrentInterests,
				ID:        req.UserID,
			})
		}

		unchangedInterests := lib.Intersection(req.PreviousInterests, req.CurrentInterests)
		toRemove := lib.Complement(unchangedInterests, req.PreviousInterests)
		toAdd := lib.Complement(unchangedInterests, req.CurrentInterests)

		if len(toAdd) > 0 {
			err := r.queue.QueueJob(ADD_TO_INTEREST_BUCKETS, AddToInterestBucketPayload{
				Interests: toAdd,
				ID:        req.UserID,
			})
			if err != nil {
				return err
			}
		}

		if len(toRemove) > 0 {
			err := r.queue.QueueJob(REMOVE_FROM_INTEREST_BUCKETS, RemoveFromInterestBucketPayload{
				Interests: toRemove,
				ID:        req.UserID,
			})
			if err != nil {
				return err
			}
		}

		return nil
	}
}
