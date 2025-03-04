package inferenceendpoints

import "encoding/json"

// Represents a list of inference endpoint jobs
type InferenceEndpointJobList struct {
	// The list of inference endpoint jobs.
	Items []InferenceEndpointJob `json:"items,omitempty" required:"true" maxItems:"100"`
	// The page number.
	Page *int64 `json:"page,omitempty" required:"true" min:"1" max:"2147483647"`
	// The maximum number of items per page.
	PageSize *int64 `json:"page_size,omitempty" required:"true" min:"1" max:"100"`
	// The total number of items in the collection.
	TotalSize *int64 `json:"total_size,omitempty" required:"true" min:"0" max:"2147483647"`
}

func (i *InferenceEndpointJobList) GetItems() []InferenceEndpointJob {
	if i == nil {
		return nil
	}
	return i.Items
}

func (i *InferenceEndpointJobList) SetItems(items []InferenceEndpointJob) {
	i.Items = items
}

func (i *InferenceEndpointJobList) GetPage() *int64 {
	if i == nil {
		return nil
	}
	return i.Page
}

func (i *InferenceEndpointJobList) SetPage(page int64) {
	i.Page = &page
}

func (i *InferenceEndpointJobList) GetPageSize() *int64 {
	if i == nil {
		return nil
	}
	return i.PageSize
}

func (i *InferenceEndpointJobList) SetPageSize(pageSize int64) {
	i.PageSize = &pageSize
}

func (i *InferenceEndpointJobList) GetTotalSize() *int64 {
	if i == nil {
		return nil
	}
	return i.TotalSize
}

func (i *InferenceEndpointJobList) SetTotalSize(totalSize int64) {
	i.TotalSize = &totalSize
}

func (i InferenceEndpointJobList) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InferenceEndpointJobList to string"
	}
	return string(jsonData)
}
