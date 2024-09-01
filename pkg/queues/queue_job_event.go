package queues

// Represents an event for queue job
type QueueJobEvent struct {
	Action *QueueJobEventAction `json:"action,omitempty" required:"true"`
	Time   *string              `json:"time,omitempty" required:"true"`
}

func (q *QueueJobEvent) SetAction(action QueueJobEventAction) {
	q.Action = &action
}

func (q *QueueJobEvent) GetAction() *QueueJobEventAction {
	if q == nil {
		return nil
	}
	return q.Action
}

func (q *QueueJobEvent) SetTime(time string) {
	q.Time = &time
}

func (q *QueueJobEvent) GetTime() *string {
	if q == nil {
		return nil
	}
	return q.Time
}

type QueueJobEventAction string

const (
	QUEUE_JOB_EVENT_ACTION_CREATED   QueueJobEventAction = "created"
	QUEUE_JOB_EVENT_ACTION_STARTED   QueueJobEventAction = "started"
	QUEUE_JOB_EVENT_ACTION_SUCCEEDED QueueJobEventAction = "succeeded"
	QUEUE_JOB_EVENT_ACTION_CANCELLED QueueJobEventAction = "cancelled"
	QUEUE_JOB_EVENT_ACTION_FAILED    QueueJobEventAction = "failed"
)
