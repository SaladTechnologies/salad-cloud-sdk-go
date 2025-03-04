package queues

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a queue job
type QueueJob struct {
	Id *string `json:"id,omitempty" required:"true"`
	// The job input. May be any valid JSON.
	Input    any                    `json:"input,omitempty" required:"true"`
	Metadata *util.Nullable[any]    `json:"metadata,omitempty"`
	Webhook  *util.Nullable[string] `json:"webhook,omitempty"`
	Status   *QueueJobStatus        `json:"status,omitempty" required:"true"`
	Events   []QueueJobEvent        `json:"events,omitempty" required:"true" maxItems:"1000"`
	// The job output. May be any valid JSON.
	Output     any     `json:"output,omitempty"`
	CreateTime *string `json:"create_time,omitempty" required:"true"`
	UpdateTime *string `json:"update_time,omitempty" required:"true"`
}

func (q *QueueJob) GetId() *string {
	if q == nil {
		return nil
	}
	return q.Id
}

func (q *QueueJob) SetId(id string) {
	q.Id = &id
}

func (q *QueueJob) GetInput() any {
	if q == nil {
		return nil
	}
	return q.Input
}

func (q *QueueJob) SetInput(input any) {
	q.Input = input
}

func (q *QueueJob) GetMetadata() *util.Nullable[any] {
	if q == nil {
		return nil
	}
	return q.Metadata
}

func (q *QueueJob) SetMetadata(metadata util.Nullable[any]) {
	q.Metadata = &metadata
}

func (q *QueueJob) SetMetadataNull() {
	q.Metadata = &util.Nullable[any]{IsNull: true}
}

func (q *QueueJob) GetWebhook() *util.Nullable[string] {
	if q == nil {
		return nil
	}
	return q.Webhook
}

func (q *QueueJob) SetWebhook(webhook util.Nullable[string]) {
	q.Webhook = &webhook
}

func (q *QueueJob) SetWebhookNull() {
	q.Webhook = &util.Nullable[string]{IsNull: true}
}

func (q *QueueJob) GetStatus() *QueueJobStatus {
	if q == nil {
		return nil
	}
	return q.Status
}

func (q *QueueJob) SetStatus(status QueueJobStatus) {
	q.Status = &status
}

func (q *QueueJob) GetEvents() []QueueJobEvent {
	if q == nil {
		return nil
	}
	return q.Events
}

func (q *QueueJob) SetEvents(events []QueueJobEvent) {
	q.Events = events
}

func (q *QueueJob) GetOutput() any {
	if q == nil {
		return nil
	}
	return q.Output
}

func (q *QueueJob) SetOutput(output any) {
	q.Output = output
}

func (q *QueueJob) GetCreateTime() *string {
	if q == nil {
		return nil
	}
	return q.CreateTime
}

func (q *QueueJob) SetCreateTime(createTime string) {
	q.CreateTime = &createTime
}

func (q *QueueJob) GetUpdateTime() *string {
	if q == nil {
		return nil
	}
	return q.UpdateTime
}

func (q *QueueJob) SetUpdateTime(updateTime string) {
	q.UpdateTime = &updateTime
}

func (q QueueJob) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueJob to string"
	}
	return string(jsonData)
}

type QueueJobStatus string

const (
	QUEUE_JOB_STATUS_PENDING   QueueJobStatus = "pending"
	QUEUE_JOB_STATUS_RUNNING   QueueJobStatus = "running"
	QUEUE_JOB_STATUS_SUCCEEDED QueueJobStatus = "succeeded"
	QUEUE_JOB_STATUS_CANCELLED QueueJobStatus = "cancelled"
	QUEUE_JOB_STATUS_FAILED    QueueJobStatus = "failed"
)
