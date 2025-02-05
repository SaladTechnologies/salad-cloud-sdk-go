package inferenceendpoints

import (
	"encoding/json"
)

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
	touched    map[string]bool
}

func (i *InferenceEndpointJob) GetId() *string {
	if i == nil {
		return nil
	}
	return i.Id
}

func (i *InferenceEndpointJob) SetId(id string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Id"] = true
	i.Id = &id
}

func (i *InferenceEndpointJob) SetIdNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Id"] = true
	i.Id = nil
}

func (i *InferenceEndpointJob) GetInput() any {
	if i == nil {
		return nil
	}
	return i.Input
}

func (i *InferenceEndpointJob) SetInput(input any) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Input"] = true
	i.Input = input
}

func (i *InferenceEndpointJob) SetInputNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Input"] = true
	i.Input = nil
}

func (i *InferenceEndpointJob) GetInferenceEndpointName() *string {
	if i == nil {
		return nil
	}
	return i.InferenceEndpointName
}

func (i *InferenceEndpointJob) SetInferenceEndpointName(inferenceEndpointName string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["InferenceEndpointName"] = true
	i.InferenceEndpointName = &inferenceEndpointName
}

func (i *InferenceEndpointJob) SetInferenceEndpointNameNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["InferenceEndpointName"] = true
	i.InferenceEndpointName = nil
}

func (i *InferenceEndpointJob) GetMetadata() any {
	if i == nil {
		return nil
	}
	return i.Metadata
}

func (i *InferenceEndpointJob) SetMetadata(metadata any) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Metadata"] = true
	i.Metadata = metadata
}

func (i *InferenceEndpointJob) SetMetadataNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Metadata"] = true
	i.Metadata = nil
}

func (i *InferenceEndpointJob) GetWebhook() *string {
	if i == nil {
		return nil
	}
	return i.Webhook
}

func (i *InferenceEndpointJob) SetWebhook(webhook string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Webhook"] = true
	i.Webhook = &webhook
}

func (i *InferenceEndpointJob) SetWebhookNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Webhook"] = true
	i.Webhook = nil
}

func (i *InferenceEndpointJob) GetStatus() *InferenceEndpointJobStatus {
	if i == nil {
		return nil
	}
	return i.Status
}

func (i *InferenceEndpointJob) SetStatus(status InferenceEndpointJobStatus) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Status"] = true
	i.Status = &status
}

func (i *InferenceEndpointJob) SetStatusNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Status"] = true
	i.Status = nil
}

func (i *InferenceEndpointJob) GetEvents() []InferenceEndpointJobEvent {
	if i == nil {
		return nil
	}
	return i.Events
}

func (i *InferenceEndpointJob) SetEvents(events []InferenceEndpointJobEvent) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Events"] = true
	i.Events = events
}

func (i *InferenceEndpointJob) SetEventsNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Events"] = true
	i.Events = nil
}

func (i *InferenceEndpointJob) GetOrganizationName() *string {
	if i == nil {
		return nil
	}
	return i.OrganizationName
}

func (i *InferenceEndpointJob) SetOrganizationName(organizationName string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["OrganizationName"] = true
	i.OrganizationName = &organizationName
}

func (i *InferenceEndpointJob) SetOrganizationNameNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["OrganizationName"] = true
	i.OrganizationName = nil
}

func (i *InferenceEndpointJob) GetOutput() any {
	if i == nil {
		return nil
	}
	return i.Output
}

func (i *InferenceEndpointJob) SetOutput(output any) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Output"] = true
	i.Output = output
}

func (i *InferenceEndpointJob) SetOutputNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Output"] = true
	i.Output = nil
}

func (i *InferenceEndpointJob) GetCreateTime() *string {
	if i == nil {
		return nil
	}
	return i.CreateTime
}

func (i *InferenceEndpointJob) SetCreateTime(createTime string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["CreateTime"] = true
	i.CreateTime = &createTime
}

func (i *InferenceEndpointJob) SetCreateTimeNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["CreateTime"] = true
	i.CreateTime = nil
}

func (i *InferenceEndpointJob) GetUpdateTime() *string {
	if i == nil {
		return nil
	}
	return i.UpdateTime
}

func (i *InferenceEndpointJob) SetUpdateTime(updateTime string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["UpdateTime"] = true
	i.UpdateTime = &updateTime
}

func (i *InferenceEndpointJob) SetUpdateTimeNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["UpdateTime"] = true
	i.UpdateTime = nil
}

func (i InferenceEndpointJob) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if i.touched["Id"] && i.Id == nil {
		data["id"] = nil
	} else if i.Id != nil {
		data["id"] = i.Id
	}

	if i.touched["Input"] && i.Input == nil {
		data["input"] = nil
	} else if i.Input != nil {
		data["input"] = i.Input
	}

	if i.touched["InferenceEndpointName"] && i.InferenceEndpointName == nil {
		data["inference_endpoint_name"] = nil
	} else if i.InferenceEndpointName != nil {
		data["inference_endpoint_name"] = i.InferenceEndpointName
	}

	if i.touched["Metadata"] && i.Metadata == nil {
		data["metadata"] = nil
	} else if i.Metadata != nil {
		data["metadata"] = i.Metadata
	}

	if i.touched["Webhook"] && i.Webhook == nil {
		data["webhook"] = nil
	} else if i.Webhook != nil {
		data["webhook"] = i.Webhook
	}

	if i.touched["Status"] && i.Status == nil {
		data["status"] = nil
	} else if i.Status != nil {
		data["status"] = i.Status
	}

	if i.touched["Events"] && i.Events == nil {
		data["events"] = nil
	} else if i.Events != nil {
		data["events"] = i.Events
	}

	if i.touched["OrganizationName"] && i.OrganizationName == nil {
		data["organization_name"] = nil
	} else if i.OrganizationName != nil {
		data["organization_name"] = i.OrganizationName
	}

	if i.touched["Output"] && i.Output == nil {
		data["output"] = nil
	} else if i.Output != nil {
		data["output"] = i.Output
	}

	if i.touched["CreateTime"] && i.CreateTime == nil {
		data["create_time"] = nil
	} else if i.CreateTime != nil {
		data["create_time"] = i.CreateTime
	}

	if i.touched["UpdateTime"] && i.UpdateTime == nil {
		data["update_time"] = nil
	} else if i.UpdateTime != nil {
		data["update_time"] = i.UpdateTime
	}

	return json.Marshal(data)
}

func (i InferenceEndpointJob) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InferenceEndpointJob to string"
	}
	return string(jsonData)
}

type InferenceEndpointJobStatus string

const (
	INFERENCE_ENDPOINT_JOB_STATUS_PENDING   InferenceEndpointJobStatus = "pending"
	INFERENCE_ENDPOINT_JOB_STATUS_RUNNING   InferenceEndpointJobStatus = "running"
	INFERENCE_ENDPOINT_JOB_STATUS_SUCCEEDED InferenceEndpointJobStatus = "succeeded"
	INFERENCE_ENDPOINT_JOB_STATUS_CANCELLED InferenceEndpointJobStatus = "cancelled"
	INFERENCE_ENDPOINT_JOB_STATUS_FAILED    InferenceEndpointJobStatus = "failed"
)
