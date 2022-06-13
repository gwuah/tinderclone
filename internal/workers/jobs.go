package workers

import "github.com/gwuah/tinderclone/internal/queue"

const (
	SEND_SMS                     queue.Job = "send_sms"
	SEND_EMAIL                   queue.Job = "send_email"
	ADD_TO_INTEREST_BUCKETS      queue.Job = "add_to_interest_buckets"
	REMOVE_FROM_INTEREST_BUCKETS queue.Job = "remove_from_interest_buckets"
)
