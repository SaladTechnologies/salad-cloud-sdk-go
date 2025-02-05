package inferenceendpoints

import (
	"encoding/json"
)

// Represents an event for inference endpoint job
type InferenceEndpointJobEvent struct {
	Action  *InferenceEndpointJobEventAction `json:"action,omitempty" required:"true"`
	Time    *string                          `json:"time,omitempty" required:"true"`
	touched map[string]bool
}

func (i *InferenceEndpointJobEvent) GetAction() *InferenceEndpointJobEventAction {
	if i == nil {
		return nil
	}
	return i.Action
}

func (i *InferenceEndpointJobEvent) SetAction(action InferenceEndpointJobEventAction) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Action"] = true
	i.Action = &action
}

func (i *InferenceEndpointJobEvent) SetActionNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Action"] = true
	i.Action = nil
}

func (i *InferenceEndpointJobEvent) GetTime() *string {
	if i == nil {
		return nil
	}
	return i.Time
}

func (i *InferenceEndpointJobEvent) SetTime(time string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Time"] = true
	i.Time = &time
}

func (i *InferenceEndpointJobEvent) SetTimeNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Time"] = true
	i.Time = nil
}

func (i InferenceEndpointJobEvent) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if i.touched["Action"] && i.Action == nil {
		data["action"] = nil
	} else if i.Action != nil {
		data["action"] = i.Action
	}

	if i.touched["Time"] && i.Time == nil {
		data["time"] = nil
	} else if i.Time != nil {
		data["time"] = i.Time
	}

	return json.Marshal(data)
}

func (i InferenceEndpointJobEvent) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InferenceEndpointJobEvent to string"
	}
	return string(jsonData)
}

type InferenceEndpointJobEventAction string

const (
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_CREATED   InferenceEndpointJobEventAction = "created"
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_STARTED   InferenceEndpointJobEventAction = "started"
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_SUCCEEDED InferenceEndpointJobEventAction = "succeeded"
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_CANCELLED InferenceEndpointJobEventAction = "cancelled"
	INFERENCE_ENDPOINT_JOB_EVENT_ACTION_FAILED    InferenceEndpointJobEventAction = "failed"
)
