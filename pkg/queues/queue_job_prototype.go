package queues

import "encoding/json"

// Represents a request to create a queue job
type QueueJobPrototype struct {
	// The job input. May be any valid JSON.
	Input any `json:"input,omitempty" required:"true"`
	// Additional metadata for the job
	Metadata any `json:"metadata,omitempty"`
	// The webhook to call when the job completes
	Webhook *string `json:"webhook,omitempty" maxLength:"2048" minLength:"1" pattern:"^.*$"`
}

func (q *QueueJobPrototype) GetInput() any {
	if q == nil {
		return nil
	}
	return q.Input
}

func (q *QueueJobPrototype) SetInput(input any) {
	q.Input = input
}

func (q *QueueJobPrototype) GetMetadata() any {
	if q == nil {
		return nil
	}
	return q.Metadata
}

func (q *QueueJobPrototype) SetMetadata(metadata any) {
	q.Metadata = &metadata
}

func (q *QueueJobPrototype) GetWebhook() *string {
	if q == nil {
		return nil
	}
	return q.Webhook
}

func (q *QueueJobPrototype) SetWebhook(webhook string) {
	q.Webhook = &webhook
}

func (q QueueJobPrototype) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueJobPrototype to string"
	}
	return string(jsonData)
}
