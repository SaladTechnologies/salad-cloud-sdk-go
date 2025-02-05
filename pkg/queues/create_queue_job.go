package queues

import (
	"encoding/json"
)

// Represents a request to create a queue job
type CreateQueueJob struct {
	// The job input. May be any valid JSON.
	Input    any     `json:"input,omitempty" required:"true"`
	Metadata any     `json:"metadata,omitempty"`
	Webhook  *string `json:"webhook,omitempty"`
	touched  map[string]bool
}

func (c *CreateQueueJob) GetInput() any {
	if c == nil {
		return nil
	}
	return c.Input
}

func (c *CreateQueueJob) SetInput(input any) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Input"] = true
	c.Input = input
}

func (c *CreateQueueJob) SetInputNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Input"] = true
	c.Input = nil
}

func (c *CreateQueueJob) GetMetadata() any {
	if c == nil {
		return nil
	}
	return c.Metadata
}

func (c *CreateQueueJob) SetMetadata(metadata any) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Metadata"] = true
	c.Metadata = metadata
}

func (c *CreateQueueJob) SetMetadataNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Metadata"] = true
	c.Metadata = nil
}

func (c *CreateQueueJob) GetWebhook() *string {
	if c == nil {
		return nil
	}
	return c.Webhook
}

func (c *CreateQueueJob) SetWebhook(webhook string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Webhook"] = true
	c.Webhook = &webhook
}

func (c *CreateQueueJob) SetWebhookNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Webhook"] = true
	c.Webhook = nil
}

func (c CreateQueueJob) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Input"] && c.Input == nil {
		data["input"] = nil
	} else if c.Input != nil {
		data["input"] = c.Input
	}

	if c.touched["Metadata"] && c.Metadata == nil {
		data["metadata"] = nil
	} else if c.Metadata != nil {
		data["metadata"] = c.Metadata
	}

	if c.touched["Webhook"] && c.Webhook == nil {
		data["webhook"] = nil
	} else if c.Webhook != nil {
		data["webhook"] = c.Webhook
	}

	return json.Marshal(data)
}

func (c CreateQueueJob) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateQueueJob to string"
	}
	return string(jsonData)
}
