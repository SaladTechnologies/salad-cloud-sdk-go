package queues

// Represents a request to create a queue job
type CreateQueueJob struct {
	// The job input. May be any valid JSON.
	Input    any     `json:"input,omitempty" required:"true"`
	Metadata any     `json:"metadata,omitempty"`
	Webhook  *string `json:"webhook,omitempty"`
}

func (c *CreateQueueJob) SetInput(input any) {
	c.Input = input
}

func (c *CreateQueueJob) GetInput() any {
	if c == nil {
		return nil
	}
	return c.Input
}

func (c *CreateQueueJob) SetMetadata(metadata any) {
	c.Metadata = metadata
}

func (c *CreateQueueJob) GetMetadata() any {
	if c == nil {
		return nil
	}
	return c.Metadata
}

func (c *CreateQueueJob) SetWebhook(webhook string) {
	c.Webhook = &webhook
}

func (c *CreateQueueJob) GetWebhook() *string {
	if c == nil {
		return nil
	}
	return c.Webhook
}
