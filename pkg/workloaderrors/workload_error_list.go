package workloaderrors

import (
	"encoding/json"
)

// Represents a list of workload errors
type WorkloadErrorList struct {
	Items   []WorkloadError `json:"items,omitempty" required:"true" maxItems:"50"`
	touched map[string]bool
}

func (w *WorkloadErrorList) GetItems() []WorkloadError {
	if w == nil {
		return nil
	}
	return w.Items
}

func (w *WorkloadErrorList) SetItems(items []WorkloadError) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Items"] = true
	w.Items = items
}

func (w *WorkloadErrorList) SetItemsNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Items"] = true
	w.Items = nil
}

func (w WorkloadErrorList) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if w.touched["Items"] && w.Items == nil {
		data["items"] = nil
	} else if w.Items != nil {
		data["items"] = w.Items
	}

	return json.Marshal(data)
}

func (w WorkloadErrorList) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkloadErrorList to string"
	}
	return string(jsonData)
}
