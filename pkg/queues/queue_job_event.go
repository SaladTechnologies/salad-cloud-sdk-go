package queues

import "encoding/json"

// Represents an event for queue job
type QueueJobEvent struct {
	Action *Action `json:"action,omitempty" required:"true"`
	Time   *string `json:"time,omitempty" required:"true"`
}

func (q *QueueJobEvent) GetAction() *Action {
	if q == nil {
		return nil
	}
	return q.Action
}

func (q *QueueJobEvent) SetAction(action Action) {
	q.Action = &action
}

func (q *QueueJobEvent) GetTime() *string {
	if q == nil {
		return nil
	}
	return q.Time
}

func (q *QueueJobEvent) SetTime(time string) {
	q.Time = &time
}

func (q QueueJobEvent) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueJobEvent to string"
	}
	return string(jsonData)
}

type Action string

const (
	ACTION_CREATED   Action = "created"
	ACTION_STARTED   Action = "started"
	ACTION_SUCCEEDED Action = "succeeded"
	ACTION_CANCELLED Action = "cancelled"
	ACTION_FAILED    Action = "failed"
)
