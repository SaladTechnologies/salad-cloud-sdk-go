package queues

// Represents a queue job
type QueueJob struct {
	Id *string `json:"id,omitempty" required:"true"`
	// The job input. May be any valid JSON.
	Input    any             `json:"input,omitempty" required:"true"`
	Metadata any             `json:"metadata,omitempty"`
	Webhook  *string         `json:"webhook,omitempty"`
	Status   *QueueJobStatus `json:"status,omitempty" required:"true"`
	Events   []QueueJobEvent `json:"events,omitempty" required:"true" maxItems:"1000"`
	// The job output. May be any valid JSON.
	Output     any     `json:"output,omitempty"`
	CreateTime *string `json:"create_time,omitempty" required:"true"`
	UpdateTime *string `json:"update_time,omitempty" required:"true"`
}

func (q *QueueJob) SetId(id string) {
	q.Id = &id
}

func (q *QueueJob) GetId() *string {
	if q == nil {
		return nil
	}
	return q.Id
}

func (q *QueueJob) SetInput(input any) {
	q.Input = input
}

func (q *QueueJob) GetInput() any {
	if q == nil {
		return nil
	}
	return q.Input
}

func (q *QueueJob) SetMetadata(metadata any) {
	q.Metadata = metadata
}

func (q *QueueJob) GetMetadata() any {
	if q == nil {
		return nil
	}
	return q.Metadata
}

func (q *QueueJob) SetWebhook(webhook string) {
	q.Webhook = &webhook
}

func (q *QueueJob) GetWebhook() *string {
	if q == nil {
		return nil
	}
	return q.Webhook
}

func (q *QueueJob) SetStatus(status QueueJobStatus) {
	q.Status = &status
}

func (q *QueueJob) GetStatus() *QueueJobStatus {
	if q == nil {
		return nil
	}
	return q.Status
}

func (q *QueueJob) SetEvents(events []QueueJobEvent) {
	q.Events = events
}

func (q *QueueJob) GetEvents() []QueueJobEvent {
	if q == nil {
		return nil
	}
	return q.Events
}

func (q *QueueJob) SetOutput(output any) {
	q.Output = output
}

func (q *QueueJob) GetOutput() any {
	if q == nil {
		return nil
	}
	return q.Output
}

func (q *QueueJob) SetCreateTime(createTime string) {
	q.CreateTime = &createTime
}

func (q *QueueJob) GetCreateTime() *string {
	if q == nil {
		return nil
	}
	return q.CreateTime
}

func (q *QueueJob) SetUpdateTime(updateTime string) {
	q.UpdateTime = &updateTime
}

func (q *QueueJob) GetUpdateTime() *string {
	if q == nil {
		return nil
	}
	return q.UpdateTime
}

type QueueJobStatus string

const (
	QUEUE_JOB_STATUS_PENDING   QueueJobStatus = "pending"
	QUEUE_JOB_STATUS_RUNNING   QueueJobStatus = "running"
	QUEUE_JOB_STATUS_SUCCEEDED QueueJobStatus = "succeeded"
	QUEUE_JOB_STATUS_CANCELLED QueueJobStatus = "cancelled"
	QUEUE_JOB_STATUS_FAILED    QueueJobStatus = "failed"
)
