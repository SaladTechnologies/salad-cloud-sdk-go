package inferenceendpoints

import "encoding/json"

// Represents a collection of inference endpoint jobs
type InferenceEndpointJobCollection struct {
	// The list of inference endpoint jobs.
	Items []InferenceEndpointJob `json:"items,omitempty" required:"true" maxItems:"100"`
	// The page number.
	Page *int64 `json:"page,omitempty" required:"true" min:"1" max:"2147483647"`
	// The maximum number of items per page.
	PageSize *int64 `json:"page_size,omitempty" required:"true" min:"1" max:"100"`
	// The total number of items in the collection.
	TotalSize *int64 `json:"total_size,omitempty" required:"true" min:"0" max:"2147483647"`
}

func (i *InferenceEndpointJobCollection) GetItems() []InferenceEndpointJob {
	if i == nil {
		return nil
	}
	return i.Items
}

func (i *InferenceEndpointJobCollection) SetItems(items []InferenceEndpointJob) {
	i.Items = items
}

func (i *InferenceEndpointJobCollection) GetPage() *int64 {
	if i == nil {
		return nil
	}
	return i.Page
}

func (i *InferenceEndpointJobCollection) SetPage(page int64) {
	i.Page = &page
}

func (i *InferenceEndpointJobCollection) GetPageSize() *int64 {
	if i == nil {
		return nil
	}
	return i.PageSize
}

func (i *InferenceEndpointJobCollection) SetPageSize(pageSize int64) {
	i.PageSize = &pageSize
}

func (i *InferenceEndpointJobCollection) GetTotalSize() *int64 {
	if i == nil {
		return nil
	}
	return i.TotalSize
}

func (i *InferenceEndpointJobCollection) SetTotalSize(totalSize int64) {
	i.TotalSize = &totalSize
}

func (i InferenceEndpointJobCollection) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InferenceEndpointJobCollection to string"
	}
	return string(jsonData)
}
