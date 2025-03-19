package inferenceendpoints

import "encoding/json"

// Represents a request to create a inference endpoint job
type InferenceEndpointJobPrototype struct {
	// The job input. May be any valid JSON.
	Input any `json:"input,omitempty" required:"true"`
	// The job metadata. May be any valid JSON.
	Metadata any `json:"metadata,omitempty"`
	// The webhook URL to which the job results will be POSTed.
	Webhook *string `json:"webhook,omitempty" maxLength:"2048" minLength:"1"`
	// The webhook URL to which the job results will be POSTed.
	WebhookUrl *string `json:"webhook_url,omitempty" maxLength:"2048" minLength:"1"`
}

func (i *InferenceEndpointJobPrototype) GetInput() any {
	if i == nil {
		return nil
	}
	return i.Input
}

func (i *InferenceEndpointJobPrototype) SetInput(input any) {
	i.Input = input
}

func (i *InferenceEndpointJobPrototype) GetMetadata() any {
	if i == nil {
		return nil
	}
	return i.Metadata
}

func (i *InferenceEndpointJobPrototype) SetMetadata(metadata any) {
	i.Metadata = &metadata
}

func (i *InferenceEndpointJobPrototype) GetWebhook() *string {
	if i == nil {
		return nil
	}
	return i.Webhook
}

func (i *InferenceEndpointJobPrototype) SetWebhook(webhook string) {
	i.Webhook = &webhook
}

func (i *InferenceEndpointJobPrototype) GetWebhookUrl() *string {
	if i == nil {
		return nil
	}
	return i.WebhookUrl
}

func (i *InferenceEndpointJobPrototype) SetWebhookUrl(webhookUrl string) {
	i.WebhookUrl = &webhookUrl
}

func (i InferenceEndpointJobPrototype) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InferenceEndpointJobPrototype to string"
	}
	return string(jsonData)
}
