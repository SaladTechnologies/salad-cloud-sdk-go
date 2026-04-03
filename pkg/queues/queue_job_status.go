package queues

// The job status
type QueueJobStatus string

const (
	QUEUE_JOB_STATUS_PENDING   QueueJobStatus = "pending"
	QUEUE_JOB_STATUS_RUNNING   QueueJobStatus = "running"
	QUEUE_JOB_STATUS_SUCCEEDED QueueJobStatus = "succeeded"
	QUEUE_JOB_STATUS_CANCELLED QueueJobStatus = "cancelled"
	QUEUE_JOB_STATUS_FAILED    QueueJobStatus = "failed"
)
