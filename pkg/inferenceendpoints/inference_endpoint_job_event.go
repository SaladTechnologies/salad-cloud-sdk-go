package inferenceendpoints

// Represents an event for inference endpoint job
type InferenceEndpointJobEvent struct {
	Action *InferenceEndpointJobEventAction `json:"action,omitempty" required:"true"`
	Time   *string                          `json:"time,omitempty" required:"true"`
}

func (i *InferenceEndpointJobEvent) SetAction(action InferenceEndpointJobEventAction) {
	i.Action = &action
}

func (i *InferenceEndpointJobEvent) GetAction() *InferenceEndpointJobEventAction {
	if i == nil {
		return nil
	}
	return i.Action
}

func (i *InferenceEndpointJobEvent) SetTime(time string) {
	i.Time = &time
}

func (i *InferenceEndpointJobEvent) GetTime() *string {
	if i == nil {
		return nil
	}
	return i.Time
}

type InferenceEndpointJobEventAction string

const (
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_CREATED   InferenceEndpointJobEventAction = "created"
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_STARTED   InferenceEndpointJobEventAction = "started"
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_SUCCEEDED InferenceEndpointJobEventAction = "succeeded"
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_CANCELLED InferenceEndpointJobEventAction = "cancelled"
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_FAILED    InferenceEndpointJobEventAction = "failed"
)
