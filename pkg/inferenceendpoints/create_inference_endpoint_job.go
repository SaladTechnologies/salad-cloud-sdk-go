package inferenceendpoints

// Represents a request to create a inference endpoint job
type CreateInferenceEndpointJob struct {
	// The job input. May be any valid JSON.
	Input    any     `json:"input,omitempty" required:"true"`
	Metadata any     `json:"metadata,omitempty"`
	Webhook  *string `json:"webhook,omitempty"`
}

func (c *CreateInferenceEndpointJob) SetInput(input any) {
	c.Input = input
}

func (c *CreateInferenceEndpointJob) GetInput() any {
	if c == nil {
		return nil
	}
	return c.Input
}

func (c *CreateInferenceEndpointJob) SetMetadata(metadata any) {
	c.Metadata = metadata
}

func (c *CreateInferenceEndpointJob) GetMetadata() any {
	if c == nil {
		return nil
	}
	return c.Metadata
}

func (c *CreateInferenceEndpointJob) SetWebhook(webhook string) {
	c.Webhook = &webhook
}

func (c *CreateInferenceEndpointJob) GetWebhook() *string {
	if c == nil {
		return nil
	}
	return c.Webhook
}
