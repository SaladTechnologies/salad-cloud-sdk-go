package inferenceendpoints

import (
	"encoding/json"
)

// Represents a request to create a inference endpoint job
type CreateInferenceEndpointJob struct {
	// The job input. May be any valid JSON.
	Input    any     `json:"input,omitempty" required:"true"`
	Metadata any     `json:"metadata,omitempty"`
	Webhook  *string `json:"webhook,omitempty" maxLength:"2000"`
	touched  map[string]bool
}

func (c *CreateInferenceEndpointJob) GetInput() any {
	if c == nil {
		return nil
	}
	return c.Input
}

func (c *CreateInferenceEndpointJob) SetInput(input any) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Input"] = true
	c.Input = input
}

func (c *CreateInferenceEndpointJob) SetInputNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Input"] = true
	c.Input = nil
}

func (c *CreateInferenceEndpointJob) GetMetadata() any {
	if c == nil {
		return nil
	}
	return c.Metadata
}

func (c *CreateInferenceEndpointJob) SetMetadata(metadata any) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Metadata"] = true
	c.Metadata = metadata
}

func (c *CreateInferenceEndpointJob) SetMetadataNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Metadata"] = true
	c.Metadata = nil
}

func (c *CreateInferenceEndpointJob) GetWebhook() *string {
	if c == nil {
		return nil
	}
	return c.Webhook
}

func (c *CreateInferenceEndpointJob) SetWebhook(webhook string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Webhook"] = true
	c.Webhook = &webhook
}

func (c *CreateInferenceEndpointJob) SetWebhookNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Webhook"] = true
	c.Webhook = nil
}

func (c CreateInferenceEndpointJob) MarshalJSON() ([]byte, error) {
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

func (c CreateInferenceEndpointJob) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateInferenceEndpointJob to string"
	}
	return string(jsonData)
}
