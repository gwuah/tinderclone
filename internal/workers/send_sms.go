package workers

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bgentry/que-go"
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/queue"
)

type SMSWorker struct {
	termii *lib.Termii
}

type SMSPayload struct {
	To  string `json:"to"`
	Sms string `json:"sms"`
}

func NewSMSWorker(termii *lib.Termii) *SMSWorker {
	return &SMSWorker{termii: termii}
}

func (s *SMSWorker) Identifier() queue.Job {
	return SEND_SMS
}

func (s *SMSWorker) Worker() que.WorkFunc {
	return func(j *que.Job) error {
		var req SMSPayload
		if err := json.Unmarshal(j.Args, &req); err != nil {
			return fmt.Errorf("unmarshal job failed. args= %s | err= %w", string(j.Args), err)
		}

		if j.ErrorCount >= 2 {
			return fmt.Errorf(fmt.Sprintf("won't retry again | %s", j.LastError.String))
		}

		response, err := s.termii.SendTextMessage(req.To, req.Sms)
		if err != nil {
			return err
		}

		if response.MessageId != "" {
			// message was sent successfully
			return nil
		}

		// anytime a job is successful, return "nil"
		// any other thing and the job will be retried

		return errors.New("something happened. retain")
	}
}
