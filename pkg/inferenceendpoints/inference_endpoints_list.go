package inferenceendpoints

import (
	"encoding/json"
)

// Represents a list of inference endpoints
type InferenceEndpointsList struct {
	// The list of items
	Items   []InferenceEndpoint `json:"items,omitempty" required:"true" maxItems:"100"`
	touched map[string]bool
}

func (i *InferenceEndpointsList) GetItems() []InferenceEndpoint {
	if i == nil {
		return nil
	}
	return i.Items
}

func (i *InferenceEndpointsList) SetItems(items []InferenceEndpoint) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Items"] = true
	i.Items = items
}

func (i *InferenceEndpointsList) SetItemsNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Items"] = true
	i.Items = nil
}

func (i InferenceEndpointsList) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if i.touched["Items"] && i.Items == nil {
		data["items"] = nil
	} else if i.Items != nil {
		data["items"] = i.Items
	}

	return json.Marshal(data)
}

func (i InferenceEndpointsList) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InferenceEndpointsList to string"
	}
	return string(jsonData)
}
