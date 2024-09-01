package inferenceendpoints

// Represents a list of inference endpoint jobs
type InferenceEndpointJobList struct {
	// The list of items
	Items []InferenceEndpointJob `json:"items,omitempty" required:"true" maxItems:"100"`
}

func (i *InferenceEndpointJobList) SetItems(items []InferenceEndpointJob) {
	i.Items = items
}

func (i *InferenceEndpointJobList) GetItems() []InferenceEndpointJob {
	if i == nil {
		return nil
	}
	return i.Items
}
