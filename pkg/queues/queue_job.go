package queues

import (
	"encoding/json"
)

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
	touched    map[string]bool
}

func (q *QueueJob) GetId() *string {
	if q == nil {
		return nil
	}
	return q.Id
}

func (q *QueueJob) SetId(id string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Id"] = true
	q.Id = &id
}

func (q *QueueJob) SetIdNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Id"] = true
	q.Id = nil
}

func (q *QueueJob) GetInput() any {
	if q == nil {
		return nil
	}
	return q.Input
}

func (q *QueueJob) SetInput(input any) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Input"] = true
	q.Input = input
}

func (q *QueueJob) SetInputNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Input"] = true
	q.Input = nil
}

func (q *QueueJob) GetMetadata() any {
	if q == nil {
		return nil
	}
	return q.Metadata
}

func (q *QueueJob) SetMetadata(metadata any) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Metadata"] = true
	q.Metadata = metadata
}

func (q *QueueJob) SetMetadataNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Metadata"] = true
	q.Metadata = nil
}

func (q *QueueJob) GetWebhook() *string {
	if q == nil {
		return nil
	}
	return q.Webhook
}

func (q *QueueJob) SetWebhook(webhook string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Webhook"] = true
	q.Webhook = &webhook
}

func (q *QueueJob) SetWebhookNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Webhook"] = true
	q.Webhook = nil
}

func (q *QueueJob) GetStatus() *QueueJobStatus {
	if q == nil {
		return nil
	}
	return q.Status
}

func (q *QueueJob) SetStatus(status QueueJobStatus) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Status"] = true
	q.Status = &status
}

func (q *QueueJob) SetStatusNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Status"] = true
	q.Status = nil
}

func (q *QueueJob) GetEvents() []QueueJobEvent {
	if q == nil {
		return nil
	}
	return q.Events
}

func (q *QueueJob) SetEvents(events []QueueJobEvent) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Events"] = true
	q.Events = events
}

func (q *QueueJob) SetEventsNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Events"] = true
	q.Events = nil
}

func (q *QueueJob) GetOutput() any {
	if q == nil {
		return nil
	}
	return q.Output
}

func (q *QueueJob) SetOutput(output any) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Output"] = true
	q.Output = output
}

func (q *QueueJob) SetOutputNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Output"] = true
	q.Output = nil
}

func (q *QueueJob) GetCreateTime() *string {
	if q == nil {
		return nil
	}
	return q.CreateTime
}

func (q *QueueJob) SetCreateTime(createTime string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["CreateTime"] = true
	q.CreateTime = &createTime
}

func (q *QueueJob) SetCreateTimeNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["CreateTime"] = true
	q.CreateTime = nil
}

func (q *QueueJob) GetUpdateTime() *string {
	if q == nil {
		return nil
	}
	return q.UpdateTime
}

func (q *QueueJob) SetUpdateTime(updateTime string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["UpdateTime"] = true
	q.UpdateTime = &updateTime
}

func (q *QueueJob) SetUpdateTimeNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["UpdateTime"] = true
	q.UpdateTime = nil
}

func (q QueueJob) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if q.touched["Id"] && q.Id == nil {
		data["id"] = nil
	} else if q.Id != nil {
		data["id"] = q.Id
	}

	if q.touched["Input"] && q.Input == nil {
		data["input"] = nil
	} else if q.Input != nil {
		data["input"] = q.Input
	}

	if q.touched["Metadata"] && q.Metadata == nil {
		data["metadata"] = nil
	} else if q.Metadata != nil {
		data["metadata"] = q.Metadata
	}

	if q.touched["Webhook"] && q.Webhook == nil {
		data["webhook"] = nil
	} else if q.Webhook != nil {
		data["webhook"] = q.Webhook
	}

	if q.touched["Status"] && q.Status == nil {
		data["status"] = nil
	} else if q.Status != nil {
		data["status"] = q.Status
	}

	if q.touched["Events"] && q.Events == nil {
		data["events"] = nil
	} else if q.Events != nil {
		data["events"] = q.Events
	}

	if q.touched["Output"] && q.Output == nil {
		data["output"] = nil
	} else if q.Output != nil {
		data["output"] = q.Output
	}

	if q.touched["CreateTime"] && q.CreateTime == nil {
		data["create_time"] = nil
	} else if q.CreateTime != nil {
		data["create_time"] = q.CreateTime
	}

	if q.touched["UpdateTime"] && q.UpdateTime == nil {
		data["update_time"] = nil
	} else if q.UpdateTime != nil {
		data["update_time"] = q.UpdateTime
	}

	return json.Marshal(data)
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
