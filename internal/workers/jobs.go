package workers

import "github.com/gwuah/tinderclone/internal/queue"

const (
	SEND_SMS            queue.Job = "send_sms"
	SEND_EMAIL          queue.Job = "send_email"
	UPDATE_REDIS_BUCKET queue.Job = "update_redis_bucket"
)
