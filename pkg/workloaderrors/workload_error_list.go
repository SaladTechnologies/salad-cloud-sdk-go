package workloaderrors

// Represents a list of workload errors
type WorkloadErrorList struct {
	Items []WorkloadError `json:"items,omitempty" required:"true" maxItems:"50"`
}

func (w *WorkloadErrorList) SetItems(items []WorkloadError) {
	w.Items = items
}

func (w *WorkloadErrorList) GetItems() []WorkloadError {
	if w == nil {
		return nil
	}
	return w.Items
}
