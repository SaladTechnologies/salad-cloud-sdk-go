package workloaderrors

import "encoding/json"

// Represents a list of workload errors
type WorkloadErrorList struct {
	// A list of workload errors
	Items []WorkloadError `json:"items,omitempty" required:"true" maxItems:"50"`
}

func (w *WorkloadErrorList) GetItems() []WorkloadError {
	if w == nil {
		return nil
	}
	return w.Items
}

func (w *WorkloadErrorList) SetItems(items []WorkloadError) {
	w.Items = items
}

func (w WorkloadErrorList) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkloadErrorList to string"
	}
	return string(jsonData)
}
