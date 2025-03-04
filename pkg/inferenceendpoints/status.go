package inferenceendpoints

// The current status.
type Status string

const (
	STATUS_PENDING   Status = "pending"
	STATUS_RUNNING   Status = "running"
	STATUS_SUCCEEDED Status = "succeeded"
	STATUS_CANCELLED Status = "cancelled"
	STATUS_FAILED    Status = "failed"
)
