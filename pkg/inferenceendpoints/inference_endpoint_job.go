package inferenceendpoints

import "encoding/json"

// Represents a inference endpoint job
type InferenceEndpointJob struct {
	// The inference endpoint job identifier.
	Id *string `json:"id,omitempty" required:"true"`
	// The inference endpoint name.
	InferenceEndpointName *string `json:"inference_endpoint_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// The organization name.
	OrganizationName *string `json:"organization_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// The job input. May be any valid JSON.
	Input any `json:"input,omitempty" required:"true"`
	// The job metadata. May be any valid JSON.
	Metadata any `json:"metadata,omitempty"`
	// The webhook URL called when the job completes.
	Webhook *string `json:"webhook,omitempty" maxLength:"2048" minLength:"1"`
	// The webhook URL called when the job completes.
	WebhookUrl *string `json:"webhook_url,omitempty" maxLength:"2048" minLength:"1"`
	// The current status.
	Status *Status `json:"status,omitempty" required:"true"`
	// The list of events.
	Events []InferenceEndpointJobEvent `json:"events,omitempty" required:"true" maxItems:"1000"`
	// The job output. May be any valid JSON.
	Output any `json:"output,omitempty"`
	// The time the job was created.
	CreateTime *string `json:"create_time,omitempty" required:"true"`
	// The time the job was last updated.
	UpdateTime *string `json:"update_time,omitempty" required:"true"`
}

func (i *InferenceEndpointJob) GetId() *string {
	if i == nil {
		return nil
	}
	return i.Id
}

func (i *InferenceEndpointJob) SetId(id string) {
	i.Id = &id
}

func (i *InferenceEndpointJob) GetInferenceEndpointName() *string {
	if i == nil {
		return nil
	}
	return i.InferenceEndpointName
}

func (i *InferenceEndpointJob) SetInferenceEndpointName(inferenceEndpointName string) {
	i.InferenceEndpointName = &inferenceEndpointName
}

func (i *InferenceEndpointJob) GetOrganizationName() *string {
	if i == nil {
		return nil
	}
	return i.OrganizationName
}

func (i *InferenceEndpointJob) SetOrganizationName(organizationName string) {
	i.OrganizationName = &organizationName
}

func (i *InferenceEndpointJob) GetInput() any {
	if i == nil {
		return nil
	}
	return i.Input
}

func (i *InferenceEndpointJob) SetInput(input any) {
	i.Input = input
}

func (i *InferenceEndpointJob) GetMetadata() any {
	if i == nil {
		return nil
	}
	return i.Metadata
}

func (i *InferenceEndpointJob) SetMetadata(metadata any) {
	i.Metadata = &metadata
}

func (i *InferenceEndpointJob) GetWebhook() *string {
	if i == nil {
		return nil
	}
	return i.Webhook
}

func (i *InferenceEndpointJob) SetWebhook(webhook string) {
	i.Webhook = &webhook
}

func (i *InferenceEndpointJob) GetWebhookUrl() *string {
	if i == nil {
		return nil
	}
	return i.WebhookUrl
}

func (i *InferenceEndpointJob) SetWebhookUrl(webhookUrl string) {
	i.WebhookUrl = &webhookUrl
}

func (i *InferenceEndpointJob) GetStatus() *Status {
	if i == nil {
		return nil
	}
	return i.Status
}

func (i *InferenceEndpointJob) SetStatus(status Status) {
	i.Status = &status
}

func (i *InferenceEndpointJob) GetEvents() []InferenceEndpointJobEvent {
	if i == nil {
		return nil
	}
	return i.Events
}

func (i *InferenceEndpointJob) SetEvents(events []InferenceEndpointJobEvent) {
	i.Events = events
}

func (i *InferenceEndpointJob) GetOutput() any {
	if i == nil {
		return nil
	}
	return i.Output
}

func (i *InferenceEndpointJob) SetOutput(output any) {
	i.Output = output
}

func (i *InferenceEndpointJob) GetCreateTime() *string {
	if i == nil {
		return nil
	}
	return i.CreateTime
}

func (i *InferenceEndpointJob) SetCreateTime(createTime string) {
	i.CreateTime = &createTime
}

func (i *InferenceEndpointJob) GetUpdateTime() *string {
	if i == nil {
		return nil
	}
	return i.UpdateTime
}

func (i *InferenceEndpointJob) SetUpdateTime(updateTime string) {
	i.UpdateTime = &updateTime
}

func (i InferenceEndpointJob) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InferenceEndpointJob to string"
	}
	return string(jsonData)
}
