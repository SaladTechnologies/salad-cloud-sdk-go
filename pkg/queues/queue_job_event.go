package queues

import (
	"encoding/json"
)

// Represents an event for queue job
type QueueJobEvent struct {
	Action  *QueueJobEventAction `json:"action,omitempty" required:"true"`
	Time    *string              `json:"time,omitempty" required:"true"`
	touched map[string]bool
}

func (q *QueueJobEvent) GetAction() *QueueJobEventAction {
	if q == nil {
		return nil
	}
	return q.Action
}

func (q *QueueJobEvent) SetAction(action QueueJobEventAction) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Action"] = true
	q.Action = &action
}

func (q *QueueJobEvent) SetActionNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Action"] = true
	q.Action = nil
}

func (q *QueueJobEvent) GetTime() *string {
	if q == nil {
		return nil
	}
	return q.Time
}

func (q *QueueJobEvent) SetTime(time string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Time"] = true
	q.Time = &time
}

func (q *QueueJobEvent) SetTimeNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Time"] = true
	q.Time = nil
}

func (q QueueJobEvent) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if q.touched["Action"] && q.Action == nil {
		data["action"] = nil
	} else if q.Action != nil {
		data["action"] = q.Action
	}

	if q.touched["Time"] && q.Time == nil {
		data["time"] = nil
	} else if q.Time != nil {
		data["time"] = q.Time
	}

	return json.Marshal(data)
}

func (q QueueJobEvent) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueJobEvent to string"
	}
	return string(jsonData)
}

type QueueJobEventAction string

const (
	QUEUE_JOB_EVENT_ACTION_CREATED   QueueJobEventAction = "created"
	QUEUE_JOB_EVENT_ACTION_STARTED   QueueJobEventAction = "started"
	QUEUE_JOB_EVENT_ACTION_SUCCEEDED QueueJobEventAction = "succeeded"
	QUEUE_JOB_EVENT_ACTION_CANCELLED QueueJobEventAction = "cancelled"
	QUEUE_JOB_EVENT_ACTION_FAILED    QueueJobEventAction = "failed"
)
