package queues

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a request to create a queue job
type CreateQueueJob struct {
	// The job input. May be any valid JSON.
	Input    any                    `json:"input,omitempty" required:"true"`
	Metadata *util.Nullable[any]    `json:"metadata,omitempty"`
	Webhook  *util.Nullable[string] `json:"webhook,omitempty"`
}

func (c *CreateQueueJob) GetInput() any {
	if c == nil {
		return nil
	}
	return c.Input
}

func (c *CreateQueueJob) SetInput(input any) {
	c.Input = input
}

func (c *CreateQueueJob) GetMetadata() *util.Nullable[any] {
	if c == nil {
		return nil
	}
	return c.Metadata
}

func (c *CreateQueueJob) SetMetadata(metadata util.Nullable[any]) {
	c.Metadata = &metadata
}

func (c *CreateQueueJob) SetMetadataNull() {
	c.Metadata = &util.Nullable[any]{IsNull: true}
}

func (c *CreateQueueJob) GetWebhook() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.Webhook
}

func (c *CreateQueueJob) SetWebhook(webhook util.Nullable[string]) {
	c.Webhook = &webhook
}

func (c *CreateQueueJob) SetWebhookNull() {
	c.Webhook = &util.Nullable[string]{IsNull: true}
}

func (c CreateQueueJob) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateQueueJob to string"
	}
	return string(jsonData)
}
