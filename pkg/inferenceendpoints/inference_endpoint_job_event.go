package inferenceendpoints

import "encoding/json"

// Represents an event for inference endpoint job
type InferenceEndpointJobEvent struct {
	// The action that was taken on the inference endpoint job.
	Action *InferenceEndpointJobEventAction `json:"action,omitempty" required:"true"`
	// The time the event occurred.
	Time *string `json:"time,omitempty" required:"true"`
}

func (i *InferenceEndpointJobEvent) GetAction() *InferenceEndpointJobEventAction {
	if i == nil {
		return nil
	}
	return i.Action
}

func (i *InferenceEndpointJobEvent) SetAction(action InferenceEndpointJobEventAction) {
	i.Action = &action
}

func (i *InferenceEndpointJobEvent) GetTime() *string {
	if i == nil {
		return nil
	}
	return i.Time
}

func (i *InferenceEndpointJobEvent) SetTime(time string) {
	i.Time = &time
}

func (i InferenceEndpointJobEvent) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InferenceEndpointJobEvent to string"
	}
	return string(jsonData)
}
