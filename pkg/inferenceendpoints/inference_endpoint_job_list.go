package inferenceendpoints

import (
	"encoding/json"
)

// Represents a list of inference endpoint jobs
type InferenceEndpointJobList struct {
	// The list of items
	Items   []InferenceEndpointJob `json:"items,omitempty" required:"true" maxItems:"100"`
	touched map[string]bool
}

func (i *InferenceEndpointJobList) GetItems() []InferenceEndpointJob {
	if i == nil {
		return nil
	}
	return i.Items
}

func (i *InferenceEndpointJobList) SetItems(items []InferenceEndpointJob) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Items"] = true
	i.Items = items
}

func (i *InferenceEndpointJobList) SetItemsNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Items"] = true
	i.Items = nil
}

func (i InferenceEndpointJobList) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if i.touched["Items"] && i.Items == nil {
		data["items"] = nil
	} else if i.Items != nil {
		data["items"] = i.Items
	}

	return json.Marshal(data)
}

func (i InferenceEndpointJobList) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InferenceEndpointJobList to string"
	}
	return string(jsonData)
}
