package inferenceendpoints

import "encoding/json"

// Represents a request to create a inference endpoint job
type CreateInferenceEndpointJob struct {
	// The job input. May be any valid JSON.
	Input      any     `json:"input,omitempty" required:"true"`
	Metadata   any     `json:"metadata,omitempty"`
	Webhook    *string `json:"webhook,omitempty" maxLength:"2048"`
	WebhookUrl *string `json:"webhook_url,omitempty" maxLength:"2048"`
}

func (c *CreateInferenceEndpointJob) GetInput() any {
	if c == nil {
		return nil
	}
	return c.Input
}

func (c *CreateInferenceEndpointJob) SetInput(input any) {
	c.Input = input
}

func (c *CreateInferenceEndpointJob) GetMetadata() any {
	if c == nil {
		return nil
	}
	return c.Metadata
}

func (c *CreateInferenceEndpointJob) SetMetadata(metadata any) {
	c.Metadata = &metadata
}

func (c *CreateInferenceEndpointJob) GetWebhook() *string {
	if c == nil {
		return nil
	}
	return c.Webhook
}

func (c *CreateInferenceEndpointJob) SetWebhook(webhook string) {
	c.Webhook = &webhook
}

func (c *CreateInferenceEndpointJob) GetWebhookUrl() *string {
	if c == nil {
		return nil
	}
	return c.WebhookUrl
}

func (c *CreateInferenceEndpointJob) SetWebhookUrl(webhookUrl string) {
	c.WebhookUrl = &webhookUrl
}

func (c CreateInferenceEndpointJob) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateInferenceEndpointJob to string"
	}
	return string(jsonData)
}
