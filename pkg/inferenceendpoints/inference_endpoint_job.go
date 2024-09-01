package inferenceendpoints

// Represents a inference endpoint job
type InferenceEndpointJob struct {
	Id *string `json:"id,omitempty" required:"true"`
	// The job input. May be any valid JSON.
	Input any `json:"input,omitempty" required:"true"`
	// The inference endpoint name
	InferenceEndpointName *string                     `json:"inference_endpoint_name,omitempty" required:"true"`
	Metadata              any                         `json:"metadata,omitempty"`
	Webhook               *string                     `json:"webhook,omitempty"`
	Status                *InferenceEndpointJobStatus `json:"status,omitempty" required:"true"`
	Events                []InferenceEndpointJobEvent `json:"events,omitempty" required:"true" maxItems:"1000"`
	// The organization name
	OrganizationName *string `json:"organization_name,omitempty" required:"true"`
	// The job output. May be any valid JSON.
	Output     any     `json:"output,omitempty"`
	CreateTime *string `json:"create_time,omitempty" required:"true"`
	UpdateTime *string `json:"update_time,omitempty" required:"true"`
}

func (i *InferenceEndpointJob) SetId(id string) {
	i.Id = &id
}

func (i *InferenceEndpointJob) GetId() *string {
	if i == nil {
		return nil
	}
	return i.Id
}

func (i *InferenceEndpointJob) SetInput(input any) {
	i.Input = input
}

func (i *InferenceEndpointJob) GetInput() any {
	if i == nil {
		return nil
	}
	return i.Input
}

func (i *InferenceEndpointJob) SetInferenceEndpointName(inferenceEndpointName string) {
	i.InferenceEndpointName = &inferenceEndpointName
}

func (i *InferenceEndpointJob) GetInferenceEndpointName() *string {
	if i == nil {
		return nil
	}
	return i.InferenceEndpointName
}

func (i *InferenceEndpointJob) SetMetadata(metadata any) {
	i.Metadata = metadata
}

func (i *InferenceEndpointJob) GetMetadata() any {
	if i == nil {
		return nil
	}
	return i.Metadata
}

func (i *InferenceEndpointJob) SetWebhook(webhook string) {
	i.Webhook = &webhook
}

func (i *InferenceEndpointJob) GetWebhook() *string {
	if i == nil {
		return nil
	}
	return i.Webhook
}

func (i *InferenceEndpointJob) SetStatus(status InferenceEndpointJobStatus) {
	i.Status = &status
}

func (i *InferenceEndpointJob) GetStatus() *InferenceEndpointJobStatus {
	if i == nil {
		return nil
	}
	return i.Status
}

func (i *InferenceEndpointJob) SetEvents(events []InferenceEndpointJobEvent) {
	i.Events = events
}

func (i *InferenceEndpointJob) GetEvents() []InferenceEndpointJobEvent {
	if i == nil {
		return nil
	}
	return i.Events
}

func (i *InferenceEndpointJob) SetOrganizationName(organizationName string) {
	i.OrganizationName = &organizationName
}

func (i *InferenceEndpointJob) GetOrganizationName() *string {
	if i == nil {
		return nil
	}
	return i.OrganizationName
}

func (i *InferenceEndpointJob) SetOutput(output any) {
	i.Output = output
}

func (i *InferenceEndpointJob) GetOutput() any {
	if i == nil {
		return nil
	}
	return i.Output
}

func (i *InferenceEndpointJob) SetCreateTime(createTime string) {
	i.CreateTime = &createTime
}

func (i *InferenceEndpointJob) GetCreateTime() *string {
	if i == nil {
		return nil
	}
	return i.CreateTime
}

func (i *InferenceEndpointJob) SetUpdateTime(updateTime string) {
	i.UpdateTime = &updateTime
}

func (i *InferenceEndpointJob) GetUpdateTime() *string {
	if i == nil {
		return nil
	}
	return i.UpdateTime
}

type InferenceEndpointJobStatus string

const (
	INFERENCE_ENDPOINT_JOB_STATUS_PENDING   InferenceEndpointJobStatus = "pending"
	INFERENCE_ENDPOINT_JOB_STATUS_RUNNING   InferenceEndpointJobStatus = "running"
	INFERENCE_ENDPOINT_JOB_STATUS_SUCCEEDED InferenceEndpointJobStatus = "succeeded"
	INFERENCE_ENDPOINT_JOB_STATUS_CANCELLED InferenceEndpointJobStatus = "cancelled"
	INFERENCE_ENDPOINT_JOB_STATUS_FAILED    InferenceEndpointJobStatus = "failed"
)
