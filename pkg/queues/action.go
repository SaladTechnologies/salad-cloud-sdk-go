package queues

// The action that was taken on the queue job
type Action string

const (
	ACTION_CREATED   Action = "created"
	ACTION_STARTED   Action = "started"
	ACTION_SUCCEEDED Action = "succeeded"
	ACTION_CANCELLED Action = "cancelled"
	ACTION_FAILED    Action = "failed"
)
