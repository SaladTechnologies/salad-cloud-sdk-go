package inferenceendpoints

// Represents a list of inference endpoints
type InferenceEndpointsList struct {
	// The list of items
	Items []InferenceEndpoint `json:"items,omitempty" required:"true" maxItems:"100"`
}

func (i *InferenceEndpointsList) SetItems(items []InferenceEndpoint) {
	i.Items = items
}

func (i *InferenceEndpointsList) GetItems() []InferenceEndpoint {
	if i == nil {
		return nil
	}
	return i.Items
}
