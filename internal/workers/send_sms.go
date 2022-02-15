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
	sms *lib.SMS
}

type SMSPayload struct {
	To  string `json:"to"`
	Sms string `json:"sms"`
}

func NewSMSWorker(sms *lib.SMS) *SMSWorker {
	return &SMSWorker{sms: sms}
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

		response, err := s.sms.SendTextMessage(req.To, req.Sms)
		if err != nil {
			return err
		}

		if response.MessageId != "" {
			// message was sent successfully
			return nil
		}

		// anytime a job is succcesful, return "nil"
		// any other thing and the job will be retried

		return errors.New("something happened. retain")
	}
}
